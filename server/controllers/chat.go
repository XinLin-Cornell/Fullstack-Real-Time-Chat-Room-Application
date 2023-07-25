package controllers

import (
	"chat/http/requests"
	"chat/utils/errors"
	"encoding/json"
	"net/http"
	"strconv"

	"chat/services"
	"chat/utils"

	"github.com/gorilla/mux"
)

type ChatController struct {
	chatService services.Chat
}

func (c *ChatController) RegisterService(s services.Chat) {
	c.chatService = s
}

func (c *ChatController) ChatRooms(w http.ResponseWriter, r *http.Request) {

	res, err := c.chatService.ChatRooms()
	if err != nil {
		utils.ErrResponse(err, w)
		return
	}

	data, err := json.Marshal(res)
	errors.ErrorCheck(err)

	utils.Ok(data, w)
}

func (c *ChatController) Create(w http.ResponseWriter, r *http.Request) {
	cP := requests.ChatRoomCreatePayload{}
	err := utils.ParseBody(r, &cP)
	if err != nil {
		utils.ErrResponse(errors.ErrInRequestMarshaling, w)
		return
	}

	res, err := c.chatService.CreateChatRoom(cP.Name)
	if err != nil {
		utils.ErrResponse(err, w)
		return
	}

	data, err := json.Marshal(res)
	errors.ErrorCheck(err)

	utils.Ok(data, w)
}

func (c *ChatController) ChatRoomMessages(w http.ResponseWriter, r *http.Request) {
	cmP := requests.ChatRoomMessagesPayload{}
	err := utils.ParseBody(r, &cmP)
	if err != nil {
		utils.ErrResponse(errors.ErrInRequestMarshaling, w)
		return
	}

	res, err := c.chatService.ChatRoomMessages(cmP.RoomId)
	if err != nil {
		utils.ErrResponse(err, w)
		return
	}

	data, err := json.Marshal(res)
	errors.ErrorCheck(err)

	utils.Ok(data, w)
}

func (c *ChatController) UpvoteMessage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	messageID, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.ErrResponse(errors.ErrInRequestMarshaling, w)
		return
	}

	res, err := c.chatService.ChatUpvote(uint(messageID))
	if err != nil {
		utils.ErrResponse(err, w)
		return
	}

	data, err := json.Marshal(res)
	errors.ErrorCheck(err)

	utils.Ok(data, w)
}

func (c *ChatController) DownvoteMessage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	messageID, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.ErrResponse(errors.ErrInRequestMarshaling, w)
		return
	}

	res, err := c.chatService.ChatDownvote(uint(messageID))
	if err != nil {
		utils.ErrResponse(err, w)
		return
	}

	data, err := json.Marshal(res)
	errors.ErrorCheck(err)

	utils.Ok(data, w)
}
