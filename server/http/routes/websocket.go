package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"chat/http/responses"
	"chat/services/websocket"
	"chat/utils/errors"

	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
)

var RegisterWebsocketRoute = func(router *mux.Router) {
	pool := websocket.NewPool()
	go pool.Start()
	sb := router.PathPrefix("/v1").Subrouter()

	sb.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		jwtToken := r.URL.Query().Get("jwt")
		jwtSecret := os.Getenv("JWT_SECRET")
		token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(jwtSecret), nil
		})

		if err != nil {
			handleWebsocketAuthenticationErr(w, err)
			return
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			handleWebsocketAuthenticationErr(w, err)
			return
		}

		serveWS(pool, w, r, claims)
	})

}

func serveWS(pool *websocket.Pool, w http.ResponseWriter, r *http.Request, claims jwt.MapClaims) {
	conn, err := websocket.Upgrade(w, r)
	errors.ErrorCheck(err)

	client := &websocket.Client{
		Connection: conn,
		Pool:       pool,
		Email:      claims["Email"].(string),
		UserName:   claims["UserName"].(string),
		UserID:     uint(claims["UserID"].(float64)),
	}

	pool.Register <- client
	requestBody := make(chan []byte) // websocket.Message byte array channel
	go client.Read(requestBody)
}

func handleWebsocketAuthenticationErr(w http.ResponseWriter, err error) {
	log.Println("websocket error: ", err)
	w.WriteHeader(http.StatusUnauthorized)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	res := responses.ErrorResponse{Message: err.Error(), Status: false, Code: http.StatusUnauthorized}
	data, err := json.Marshal(res)
	errors.ErrorCheck(err)
	w.Write(data)
}
