package controllers

import (
	"fmt"
	"net/http"
	"real-time-forum/services/notification/database"
	"real-time-forum/services/notification/models"
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

	Clients = make(map[int]*websocket.Conn, 0)
	Disconnect_channel = make(chan int)
	Connection_channel = make(chan int)
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
	CustomRoute := r.Context().Value("CustomRoute").(map[string]string)

	conn, err := upgrad.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "Could not upgrade to WebSocket connection", http.StatusBadRequest)
		return
	}

	receiverId, _ := strconv.Atoi(CustomRoute["receiverId"])
	Clients[receiverId] = conn
	Connection_channel <- receiverId
	go HandleNotification()
	go CheckDisconnect(conn)
}

func HandleNotification() {
	for {

		notif := <-userNotif
		fmt.Println("Sending notification to user: ", notif.Id)

		for id, client := range Clients {
			if id == notif.ReceiverId {
				user_infos := models.UserInfos{
					Type: "notification",
					Id: notif.SenderId,
					Status: "",
				}
				err := client.WriteJSON(user_infos)
				if err == nil {
					database.DbNotification.Storage.SetModel("Id", notif.Id, models.UserNotification{}).
						UpdateField("true", "Read").Update(database.DbNotification.Storage.Db)
				}
			}
		}
	}
}

func CheckDisconnect(conn *websocket.Conn) {
	for {
		defer func() {
			disconnectUser(conn)
			conn.Close()
		}()
		_, _, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Error reading message: ", err)
			break
		}

	}
}

func disconnectUser(conn *websocket.Conn) {
	fmt.Println("Disconnecting user")
	for key, value := range Clients {
		if value == conn {
			delete(Clients, key)
			Disconnect_channel <- key
			break
		}
	}
}
