package controllers

import (
	"encoding/json"
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

	Clients        = make(map[int]*websocket.Conn, 0)
	sendMessage    = make(chan []byte)
	Broadcast      = make(chan models.Notification)
	TypingProgress = make(chan models.Typing)
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

// func (sm *SendMessage) sendMessage(w http.ResponseWriter, r *http.Request) {
// 	var response models.Response
// 	CustomRoute := r.Context().Value("CustomRoute").(map[string]string)
// 	userID := r.URL.Query().Get("user_id")

// 	if userID == "" {
// 		response.Message = "Missing user_id"
// 		utils.ResponseWithJSON(w, response, http.StatusBadRequest)
// 		return
// 	}

// 	conn, err := upgrad.Upgrade(w, r, nil)
// 	if err != nil {
// 		response.Message = "Could not upgrade to WebSocket connection"
// 		utils.ResponseWithJSON(w, "Could not upgrade to WebSocket connection", http.StatusBadRequest)
// 		return
// 	}

// 	userID_int, _ := strconv.Atoi(userID)
// 	receiverID, _ := strconv.Atoi(CustomRoute["receiverId"])
// 	Clients[userID_int] = conn

// 	// go func() {
// 	// 	defer func() {
// 	// 		conn.Close()
// 	// 		delete(Clients, userID_int)
// 	// 	}()

// 	// 	for {

// 	// 	}
// 	// }()

// 	go func(sender, receiver int) {
// 		defer func() {
// 			conn.Close()
// 			delete(Clients, sender)
// 		}()

// 		fmt.Println("sender and receiver: ", sender, receiver)
// 		for {
// 			_, msg, err := conn.ReadMessage()
// 			if err != nil {
// 				fmt.Printf("Error reading message: %s\n", err)
// 				break
// 			}

// 			var typingInfos models.Typing
// 			err = json.Unmarshal(msg, &typingInfos)

// 			if err == nil {
// 				fmt.Println("is typing...")
// 				typingInfos.Id = receiver
// 				TypingProgress <- typingInfos
// 				continue
// 			}

// 			database.DbChat.Storage.Insert(models.NewMessage(sender, receiver, string(msg)))
// 			sendMessage <- msg
// 		}
// 	}(userID_int, receiverID)


// 	go func(sender, receiver int) {
// 		defer func() {
// 			conn.Close()
// 			delete(Clients, sender)
// 		}()

// 		fmt.Println("sender and receiver: ", sender, receiver)
// 		for {
// 			msg := <-sendMessage
// 			client, ok := Clients[receiver]

// 			if ok {
// 				if err := client.WriteMessage(websocket.TextMessage, msg); err != nil {
// 					fmt.Printf("Error writing message: %s\n", err)
// 				}
// 			}
			
// 			go func() {
// 				fmt.Println("Sending message to broadcast")
// 				Broadcast <- models.NewNotification(sender, receiver, "false")
// 			}()
// 		}
// 	}(userID_int, receiverID)
// 	// for {
// 	// 	defer func() {
// 	// 		conn.Close()
// 	// 	}()
// 	// 	_, msg, err := conn.ReadMessage()
// 	// 	if err != nil {
// 	// 		fmt.Printf("Error reading message: %s\n", err)
// 	// 		break
// 	// 	}

// 	// 	idU, _ := strconv.Atoi(userID)
// 	// 	idReceiver, _ := strconv.Atoi(CustomRoute["receiverId"])

// 	// 	var typingInfos models.Typing

// 	// 	err = json.Unmarshal(msg, &typingInfos)

// 	// 	if err == nil {
// 	// 		fmt.Println("is typing...")
// 	// 		typingInfos.Id = idReceiver
// 	// 		TypingProgress <- typingInfos
// 	// 		continue
// 	// 	}

// 	// 	client, ok := Clients[idReceiver]
// 	// 	database.DbChat.Storage.Insert(models.NewMessage(idU, idReceiver, string(msg)))

// 	// 	database.DbChat.Storage.Custom.OrderBy("Id", 1).Limit(1)
// 	// 	message := database.DbChat.Storage.Scan(models.Message{}, "Id").([]models.Message)[0]
// 	// 	database.DbChat.Storage.Custom.Clear()

// 	// 	if ok {
// 	// 		if err := client.WriteMessage(websocket.TextMessage, []byte(msg)); err != nil {
// 	// 			fmt.Printf("Error writing message: %s\n", err)
// 	// 		}
// 	// 	}

// 	// 	go func(message models.Message) {
// 	// 		fmt.Println("Sending message to broadcast")
// 	// 		Broadcast <- models.NewNotification(idU, idReceiver, "false")
// 	// 	}(message)
// 	// }
// }


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

	// Channel pour envoyer les messages
	sendChannel := make(chan []byte)

	// Goroutine pour écrire sur channel de typing et indiqué à l'utilisateur que son interlocuteur
	// n'est pas entrain d'écrire
	go func() {
		for {
			
		}
	}()

	// Goroutine pour recevoir les messages
	go func() {
		defer func() {
			conn.Close()
			delete(Clients, userID_int)
		}()

		for {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				fmt.Printf("Error reading message: %s\n", err)
				break
			}

			idU, _ := strconv.Atoi(userID)
			idReceiver, _ := strconv.Atoi(CustomRoute["receiverId"])

			var typingInfos models.Typing
			err = json.Unmarshal(msg, &typingInfos)

			if err == nil {
				fmt.Println("is typing...")
				typingInfos.Id = idReceiver
				TypingProgress <- typingInfos
				continue
			}

			database.DbChat.Storage.Insert(models.NewMessage(idU, idReceiver, string(msg)))
			sendChannel <- msg
		}
	}()

	// Goroutine pour envoyer les messages
	go func() {
		defer func() {
			conn.Close()
			delete(Clients, userID_int)
		}()

		for msg := range sendChannel {
			idU, _ := strconv.Atoi(userID)
			idReceiver, _ := strconv.Atoi(CustomRoute["receiverId"])

			client, ok := Clients[idReceiver]
			if ok {
				if err := client.WriteMessage(websocket.TextMessage, msg); err != nil {
					fmt.Printf("Error writing message: %s\n", err)
				}
			}

			// Diffuser la notification
			go func() {
				fmt.Println("Sending message to broadcast")
				Broadcast <- models.NewNotification(idU, idReceiver, "false")
			}()
		}
	}()

	// Attendre que les goroutines se terminent
	select {}
}

