package chat

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"real-time-forum/server/microservices"
	"real-time-forum/server/middleware"
	"real-time-forum/services/chat/controllers"
	"real-time-forum/services/chat/database"
	"real-time-forum/services/chat/models"
	"real-time-forum/utils"
	"sync"

	"github.com/gorilla/websocket"
)

const (
	DB_NAME = "chat.db"
	DB_PATH = "../../services/chat/database/"
)

type Chat struct {
	Chat *microservices.Microservice
}

func (chat *Chat) ConfigureEndpoint() {
	for _, controller := range chat.Chat.Controllers {
		chat.Chat.Router.Method(http.MethodGet).
			Middleware(middleware.LogRequest).
			Handler(controller.EndPoint(), controller.HTTPServe())
	}
}

func (chat *Chat) InitService() (err error) {
	database.DbChat.Storage, err = utils.InitStorage(DB_NAME, DB_PATH, models.Message{})
	controllers := []microservices.Controller{
		// add controller ...
		&controllers.SendMessage{},
		&controllers.GetPrivateMessage{},
		&controllers.GetPrivateMessageUsers{},
	}

	chat.Chat = microservices.NewMicroservice("Realtime Chat", ":9090")
	chat.Chat.Controllers = append(chat.Chat.Controllers, controllers...)
	go chat.HandleNotification()
	go chat.HandleTyping()
	return
}

func (chat *Chat) GetService() *microservices.Microservice {
	return chat.Chat
}

func (chat *Chat) HandleNotification() {
	var mutex sync.Mutex
	for {

		mutex.Lock()
		notif := <-controllers.Broadcast
		mutex.Unlock()

		chat.Chat.Client.SetMethod(http.MethodPost)
		err := utils.LoadEnv("../../.env")
		if err != nil {
			log.Println(err)
		}
		baseUrl := os.Getenv("NOTIFICATION_SERVICE")
		chat.Chat.Client.SetBaseURL(baseUrl[1 : len(baseUrl)-1])

		var R models.Response
		err = chat.Chat.Client.Call("notification", "createNotification", notif, &R)
		if err != nil {
			log.Println(err)
		}
	}
}

func (chat *Chat) HandleTyping() {
	var mutex sync.Mutex
	for {
		mutex.Lock()
		typing := <-controllers.TypingProgress
		mutex.Unlock()

		conn, ok := controllers.Clients[typing.Id]

		if !ok {
			continue
		}

		data, err := json.Marshal(typing)
		if err != nil {
			log.Println("Error marshalling data: ", err)
			break
		}
		if err := conn.WriteMessage(websocket.TextMessage, data); err != nil {
			log.Println("Error writing message: ", err)
		}

	}
}
