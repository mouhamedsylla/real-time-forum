package controllers

import (
	"fmt"
	"net/http"
	"real-time-forum/services/chat/database"
	"real-time-forum/services/chat/models"
	"real-time-forum/utils"
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

	Clients   = make(map[int]*websocket.Conn, 0)
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
	var response models.Response
	CustomRoute := r.Context().Value("CustomRoute").(map[string]string)
	userID := r.URL.Query().Get("user_id")

	if userID == "" {
		response.Message = "Missing user_id"
		utils.ResponseWithJSON(w, response, http.StatusBadRequest)
		return
	}

	conn, err := upgrad.Upgrade(w, r, nil)
	if err != nil {
		response.Message = "Could not upgrade to WebSocket connection"
		utils.ResponseWithJSON(w, "Could not upgrade to WebSocket connection", http.StatusBadRequest)
		return
	}

	userID_int, _ := strconv.Atoi(userID)
	Clients[userID_int] = conn
	for {
		defer func ()  {
			conn.Close()	
		}()
		_, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Printf("Error reading message: %s\n", err)
			break
		}
		

		idU, _ := strconv.Atoi(userID)
		idReceiver, _ := strconv.Atoi(CustomRoute["receiverId"])
		client, ok := Clients[idReceiver]
		database.DbChat.Storage.Insert(models.NewMessage(idU, idReceiver, string(msg)))

		database.DbChat.Storage.Custom.OrderBy("Id", 1).Limit(1)
		message := database.DbChat.Storage.Scan(models.Message{}, "Id").([]models.Message)[0]
		database.DbChat.Storage.Custom.Clear()

		if ok {
			if err := client.WriteMessage(websocket.TextMessage, []byte(msg)); err != nil {
				fmt.Printf("Error writing message: %s\n", err)
			}
		}

		go func(message models.Message) {
			Broadcast <- models.NewNotification(idU, idReceiver, "false")
		}(message)
	}
}
