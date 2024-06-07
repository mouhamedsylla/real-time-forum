package notification

import (
	"encoding/json"
	"log"
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

func (sn *SendNotification) HTTPServe() http.Handler {
	return http.HandlerFunc(sn.SendNotification)
}

func (sn *SendNotification) EndPoint() string {
	return "/notification/sendNotification/:receiverId"
}

func (sn *SendNotification) SetMethods() []string {
	return []string{"GET"}
}

func (sn *SendNotification) SendNotification(w http.ResponseWriter, r *http.Request) {
	// Handle request logic here
	CustomRoute := r.Context().Value("CustomRoute").(map[string]string)

	conn, err := upgrad.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "Could not upgrade to WebSocket connection", http.StatusBadRequest)
		return
	}

	clients[CustomRoute["receiverId"]] = conn
	go HandleNotification(conn)
}


func HandleNotification(conn *websocket.Conn) {
	for {
		// Handle notification logic here
		notif := <- userNotif
		data, err := json.Marshal(notif)
		if err != nil {
			log.Println(err)
		}

		conn.WriteMessage(websocket.TextMessage, []byte(data))
	}
}
