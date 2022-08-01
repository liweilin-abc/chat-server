package client

import (
	"net"

	"github.com/cc-chat/message"
	"github.com/cc-chat/user"
)

type Client struct {
	Id      string               `json:"id"`
	User    user.User            `json:"user"`
	Conn    *net.Conn            `json:"_"`
	Message chan message.Message `json:"message"`
}
