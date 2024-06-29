package controllers

import (
	"fmt"
	"net/http"
	"real-time-forum/services/chat/database"
	"real-time-forum/services/chat/models"
	"strconv"

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

	clients   = make(map[string]*websocket.Conn, 0)
	Broadcast = make(chan models.Notification)
)

func (sm *SendMessage) HTTPServe() http.Handler {
	return http.HandlerFunc(sm.sendMessage)
}

func (sm *SendMessage) EndPoint() string {
	return "/chat/message/private/send/:receiverId"
}

func (sm *SendMessage) SetMethods() []string {
	return []string{http.MethodGet}
}

func (sm *SendMessage) sendMessage(w http.ResponseWriter, r *http.Request) {
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
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Printf("Error reading message: %s\n", err)
			break
		}
		client := clients[CustomRoute["receiverId"]]
		idU, _ := strconv.Atoi(userID)
		idReceiver, _ := strconv.Atoi(CustomRoute["receiverId"])
		database.DbChat.Storage.Insert(models.NewMessage(idU, idReceiver, string(msg)))

		database.DbChat.Storage.Custom.OrderBy("Id", 1).Limit(1)
		message := database.DbChat.Storage.Scan(models.Message{}, "Id").([]models.Message)[0]
		database.DbChat.Storage.Custom.Clear()

		
		fmt.Println("Message received: ", string(msg))
		if err := client.WriteMessage(websocket.TextMessage, []byte(msg)); err != nil {
			fmt.Printf("Error writing message: %s\n", err)
		}
		
		go func(message models.Message) {
			Broadcast <- models.NewNotification(message.Id, CustomRoute["receiverId"], "false")
		}(message)
	}
}
