package server

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"strings"
	"time"

	"github.com/cc-chat/client"
	"github.com/cc-chat/config"
	"github.com/cc-chat/message"
	"github.com/cc-chat/user"
	log "github.com/sirupsen/logrus"
)

type Server struct {
	Clients  map[*client.Client]bool
	Online   chan *client.Client
	Leaving  chan *client.Client
	Messages chan message.Message
	LogFile  *os.File
}

func NewServer() (*Server, error) {
	log.Printf("Opening message log file %q\n", config.Config.LogFile)
	f, err := os.OpenFile(config.Config.LogFile, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return nil, err
	}
	return &Server{
		Clients:  make(map[*client.Client]bool),
		Online:   make(chan *client.Client),
		Leaving:  make(chan *client.Client),
		Messages: make(chan message.Message),
		LogFile:  f,
	}, nil
}

func (cs *Server) ListenAndServe(addr string) error {

	if "" == addr {
		addr = config.Config.BindAddr + ":" + config.Config.BindPort
	}

	listener, err := net.Listen("tcp", addr)
	log.Println("Telnet Server listening on", addr)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	defer listener.Close()
	go cs.broadcast()
	for {
		conn, err := listener.Accept()

		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		go cs.handleConn(conn)
	}

}

func (cs *Server) broadcast() {
	for {
		select {
		case msg := <-cs.Messages:
			cs.logMessage(msg)
			for cli := range cs.Clients {
				cli.Message <- msg
			}
		case cli := <-cs.Online:
			cs.Clients[cli] = true
		case cli := <-cs.Leaving:
			delete(cs.Clients, cli)
		}
	}
}

func (cs *Server) handleConn(conn net.Conn) {
	defer conn.Close()
	ch := make(chan message.Message)
	userId := conn.RemoteAddr().String()
	name, err := readInput(conn, "Please Enter Name: ")
	if err != nil {
		panic(err)
	}
	user := user.User{
		Name: name,
	}
	// msg := message.Message{}
	// fmt.Print(user)
	client := &client.Client{userId, user, &conn, ch}

	go cs.writeLoop(client)
	ch <- message.Message{Text: "welcome " + client.User.Name, Sender: &user, Date: time.Now(), MsgType: message.SYSTEM}

	cs.Messages <- message.Message{Text: client.User.Name + " join chat", Date: time.Now(), MsgType: message.SYSTEM}
	cs.Online <- client
	cs.readLoop(client)
	cs.Messages <- message.Message{Text: client.User.Name + " has left", Date: time.Now(), MsgType: message.SYSTEM}
	cs.Leaving <- client
}

func (cs *Server) writeLoop(client *client.Client) {
	for msg := range client.Message {
		fmt.Fprintf(*client.Conn, "%s\n", msg.Text+" "+msg.Date.Format("2006-01-02 15:04:05"))
	}
}

func (cs *Server) readLoop(client *client.Client) {
	input := bufio.NewScanner(*client.Conn)
	for input.Scan() {
		cs.Messages <- message.Message{Text: client.User.Name + ": " + input.Text(), MsgType: message.USER, Date: time.Now(), Sender: &client.User}
	}
}

func readInput(conn net.Conn, msg string) (string, error) {
	conn.Write([]byte(msg))
	s, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		log.Printf("readinput: could not read input from stdin: %v from client %v", err, conn.RemoteAddr().String())
		return "", err
	}
	s = strings.Trim(s, "\r\n")
	return s, nil
}

func (cs *Server) logMessage(msg message.Message) {
	logString, err := json.Marshal(msg)
	if err != nil {
		log.Printf("Failed to Parse Message %s ", err.Error())
	}
	_, err = cs.LogFile.WriteString(string(logString) + "\n")
	if err != nil {
		log.Printf("Failed to log message%s ", err.Error())
	}
}
