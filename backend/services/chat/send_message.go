package chat

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var (
	upgrad = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	clients = make(map[string]*websocket.Conn, 0)
)

func (sm *sendMessage) HTTPServe() http.Handler {
	return http.HandlerFunc(sm.sendMessage)
}

func (sm *sendMessage) EndPoint() string {
	return "/api/message/private/send/:receiverId"
}

func (sm *sendMessage) sendMessage(w http.ResponseWriter, r *http.Request) {
	CustomRoute := r.Context().Value("CustomRoute").(map[string]string)

	userID := r.URL.Query().Get("user_id")
	if userID == "" {
		http.Error(w, "Missing user_id", http.StatusBadRequest)
		return
	}

	conn, err := upgrad.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "Could not upgrade to WebSocket connection", http.StatusBadRequest)
		return
	}

	clients[userID] = conn

	//fmt.Println("receiver: ", CustomRoute["receiverId"], " & sender: ", userID)
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Printf("Error reading message: %s\n", err)
			break
		}

		client := clients[CustomRoute["receiverId"]]
		if err := client.WriteMessage(websocket.TextMessage, []byte(msg)); err != nil {
			fmt.Printf("Error writing message: %s\n", err)
		}
	}
}

