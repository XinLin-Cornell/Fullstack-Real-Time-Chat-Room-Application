package routes

import (
	"chat/controllers"
	"chat/http/middlewares"
	"chat/services"

	"github.com/gorilla/mux"
)

var RegisterChatRoutes = func(router *mux.Router) {

	sb := router.PathPrefix("/v1/api/chat").Subrouter()
	sb.Use(middlewares.HeaderMiddleware)
	sb.Use(middlewares.Authenticated)

	var chat controllers.ChatController
	chat.RegisterService(services.NewChatService())

	sb.HandleFunc("/create", chat.Create).Methods("POST")
	sb.HandleFunc("/rooms", chat.ChatRooms).Methods("POST")
	sb.HandleFunc("/room-messages", chat.ChatRoomMessages).Methods("POST")

	sb.HandleFunc("/messages/{id}/upvote", chat.UpvoteMessage).Methods("PUT")
	sb.HandleFunc("/messages/{id}/downvote", chat.DownvoteMessage).Methods("PUT")
}
