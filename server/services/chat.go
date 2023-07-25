package services

import (
	"chat/http/responses"
	"chat/models"
	"chat/utils/errors"
)

type Chat interface {
	CreateChatRoom(name string) (responses.ChatRoomResponse, error)
	ChatRooms() (responses.ChatRoomsResponse, error)
	ChatRoomMessages(roomId uint) (responses.ChatRoomMessagesResponse, error)
	ChatUpvote(messageID uint) (responses.ChatUpvotesResponse, error)
	ChatDownvote(messageID uint) (responses.ChatDownvotesResponse, error)
}

type ChatSaver interface {
	SaveChatMessage(msg string, roomId, userId uint) uint
}

type chat struct{}

func NewChatService() *chat {
	return &chat{}
}

func (c *chat) ChatRooms() (responses.ChatRoomsResponse, error) {
	var chtList []models.ChatRoom
	var chtRoomModel models.ChatRoom
	err := chtRoomModel.List(&chtList)
	errors.DBErrorCheck(err)
	return responses.ChatRoomsResponse{ChatRooms: chtList}, nil
}

func (c *chat) CreateChatRoom(name string) (responses.ChatRoomResponse, error) {
	cht := models.ChatRoom{
		Name: name,
	}
	err := cht.Add()
	errors.DBErrorCheck(err)
	return responses.ChatRoomResponse{ChatRoom: cht}, nil
}

func (c *chat) ChatRoomMessages(roomId uint) (responses.ChatRoomMessagesResponse, error) {
	var chtList []models.Chat
	chtMsgList := []responses.ChatMessage{}
	var chtModel models.Chat
	err := chtModel.List(roomId, &chtList)
	errors.DBErrorCheck(err)
	// transform chats
	for _, v := range chtList {
		chtMsgList = append(chtMsgList, responses.ChatMessage{
			ChatMessage:   v.Message,
			ChatUserEmail: v.User.Email,
			ChatUserName:  v.User.UserName,
			ChatRoomId:    v.ChatRoomId,
			ChatRoomName:  v.ChatRoom.Name,
			Upvotes:       v.Upvotes,
			Downvotes:     v.Downvotes,
			MessageID:     v.ID,
		})
	}
	return responses.ChatRoomMessagesResponse{Chats: chtMsgList}, nil
}

func (c *chat) SaveChatMessage(msg string, roomId, userId uint) uint {
	cht := models.Chat{
		Message:    msg,
		UserId:     userId,
		ChatRoomId: roomId,
		Upvotes:    0,
		Downvotes:  0,
	}
	err := cht.Add()
	errors.DBErrorCheck(err)

	return cht.ID
}

func (c *chat) ChatUpvote(messageId uint) (responses.ChatUpvotesResponse, error) {
	var chat models.Chat
	chat.UpdateMessageUpvotes(messageId)
	return responses.ChatUpvotesResponse{Upvotes: chat.Upvotes}, nil
}

func (c *chat) ChatDownvote(messageId uint) (responses.ChatDownvotesResponse, error) {
	var chat models.Chat
	chat.UpdateMessageDownvotes(messageId)
	return responses.ChatDownvotesResponse{Downvotes: chat.Downvotes}, nil
}
