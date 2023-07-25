package websocket

import (
	"encoding/json"
	"log"

	"chat/services"
	"chat/utils/errors"

	"github.com/gorilla/websocket"
)

type Client struct {
	ID         string
	Connection *websocket.Conn
	Pool       *Pool
	Email      string
	UserName   string
	UserID     uint
}

type Message struct {
	Type int `json:"Type,omitempty"`
	Body Body
}
type Body struct {
	ChatRoomName  string `json:"chatRoomName,omitempty"`
	ChatRoomId    int32  `json:"chatRoomId,omitempty"`
	ChatMessage   string `json:"chatMessage,omitempty"`
	ChatUserEmail string `json:"chatUserEmail,omitempty"`
	ChatUserName  string `json:"chatUserName,omitempty"`
	Upvotes       uint   `json:"Upvotes,omitempty"`
	Downvotes     uint   `json:"Downvotes,omitempty"`
	MessageID     uint   `json:"MessageID,omitempty"`
}

func (c *Client) Read(bodyChan chan []byte) {
	defer func() {
		c.Pool.Unregister <- c
		c.Connection.Close()
	}()
	defer c.Pool.ReviveWebsocket()

	for {
		messageType, p, err := c.Connection.ReadMessage()
		errors.ErrorCheck(err)
		var body Body
		err = json.Unmarshal(p, &body)
		errors.ErrorCheck(err)
		body.ChatUserEmail = c.Email
		body.ChatUserName = c.UserName
		body.Upvotes = 0
		body.Downvotes = 0
		var chatSaver services.ChatSaver = services.NewChatService()
		messageID := chatSaver.SaveChatMessage(body.ChatMessage, uint(body.ChatRoomId), c.UserID)
		body.MessageID = messageID

		message := Message{Type: messageType, Body: body}
		c.Pool.Broadcast <- message
		log.Println("info:", "Message received: ", body, "messageType: ", messageType)
	}
}
