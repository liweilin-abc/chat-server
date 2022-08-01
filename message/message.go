package message

import (
	"time"

	"github.com/cc-chat/user"
)

type MsgType string

const (
	SYSTEM MsgType = "system"
	USER   MsgType = "user"
)

type Message struct {
	Id      string     `json:"id"`
	Date    time.Time  `json:"date"`   // date/time when message was sent
	Text    string     `json:"text"`   // content of the message
	Sender  *user.User `json:"sender"` // user which sent message
	MsgType MsgType    `json:"msg_type"`
}
