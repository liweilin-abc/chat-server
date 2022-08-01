# GO-Chat!

### Telnet chat server in Golang


## List of all planned and existing functionalities
- [X] Telnet server:
    - [X] Connect a client to the server
    - [X] The server relays messages to all connected clients, including a timestamp and name of the client sending the message
- [X] Chat:
    - [X] Send a message to the server
    - [X] Log messages to a file
    - [X] Read in config from a config file for port, IP, and log file location
    - [ ] Support multiple channels or rooms
    - [ ] An HTTP API to post messages
    - [ ] An HTTP API to query for messages
    - [ ] Support changing clients changing their names


## List of used libraries
* [logrus](https://github.com/sirupsen/logrus) for logging

# Usage

Connect to the server by using a telnet client

1. start server
```bash
# go to project root folder
$ cd chat-server

# get all dependencies
$ go get ./...

# run server
$ go run main.go

# you will saw 
INFO[0000] Loading configuration file: etc/config.json  
INFO[0000] Starting Telnet Chat server                  
INFO[0000] Opening message log file "log/chat.log"      
INFO[0000] Telnet Server listening on 127.0.0.1:7001  

```

2. connect to server
Open terminal

```bash 
telnet 127.0.0.1 7001

# you will see 

Trying 127.0.0.1...
Connected to localhost.
Escape character is '^]'.
Please Enter Name:

# type name 
Please Enter Name: John
welcome John 2022-08-01 04:09:58

```
![][/screenshot/chat_1.jpg]

![][/screenshot/chat_2.jpg]