package responses

import "chat/models"

type ChatRoomsResponse struct {
	ChatRooms []models.ChatRoom `json:"ChatRooms"`
}
type ChatRoomResponse struct {
	ChatRoom models.ChatRoom `json:"ChatRoom"`
}

type ChatMessage struct {
	ChatMessage   string `json:"chatMessage"`
	ChatUserEmail string `json:"chatUserEmail"`
	ChatUserName  string `json:"chatUserName"`
	ChatRoomId    uint   `json:"chatRoomId"`
	ChatRoomName  string `json:"chatRoomName"`
	Upvotes       uint   `json:"Upvotes"`
	Downvotes     uint   `json:"Downvotes"`
	MessageID     uint   `json:"MessageID"`
}

type ChatRoomMessagesResponse struct {
	Chats []ChatMessage `json:"Chats"`
}

type ChatUpvotesResponse struct {
	Upvotes uint `json:"Upvotes"`
}

type ChatDownvotesResponse struct {
	Downvotes uint `json:"Downvotes"`
}
