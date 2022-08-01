package main

import (
	"flag"

	"github.com/cc-chat/config"
	"github.com/cc-chat/server"
	log "github.com/sirupsen/logrus"
)

func init() {
	configFile := flag.String("config", "etc/config.json", "configuration file to parse")
	flag.Parse()

	log.Println("Loading configuration file:", *configFile)
	if err := config.ReadConfiguration(*configFile); err != nil {
		log.Fatalf("Failed to parse configuration file %q: %s", *configFile, err.Error())
	}
}
func main() {
	// e := echo.New()
	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())

	// e.Logger.Fatal(e.Start(":8233"))

	// exitChan := make(chan int)
	log.Println("Starting Telnet Chat server")
	server, err := server.NewServer()
	if err != nil {
		log.Info(err)
	}
	server.ListenAndServe("127.0.0.1:7001")
}
