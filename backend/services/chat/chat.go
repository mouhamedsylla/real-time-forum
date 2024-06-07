package chat

import (
	"log"
	"net/http"
	"os"
	"real-time-forum/orm"
	"real-time-forum/server/microservices"
	"real-time-forum/server/middleware"
	"real-time-forum/utils"
	"sync"
)

const (
	DB_NAME = "chat.db"
	DB_PATH = "../../services/chat/database/"
)

var storage *orm.ORM

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
	storage, err = utils.InitStorage(DB_NAME, DB_PATH, Message{})
	controllers := []microservices.Controller{
		// add controller ...
		&sendMessage{},
		&getPrivateMessage{},
		&getPrivateMessageUsers{},
	}

	chat.Chat = microservices.NewMicroservice("Realtime Chat", ":9090")
	chat.Chat.Controllers = append(chat.Chat.Controllers, controllers...)
	go chat.HandleNotification()
	return
}

func (chat *Chat) GetService() *microservices.Microservice {
	return chat.Chat
}

func (chat *Chat) HandleNotification() {
	var mutex sync.Mutex
	for {
		notif := <-broadcast
		mutex.Lock()

		chat.Chat.Client.SetMethod(http.MethodPost)
		err := utils.LoadEnv("../../.env")
		if err != nil {
			log.Println(err)
		}
		baseUrl := os.Getenv("NOTIFICATION_SERVICE")
		chat.Chat.Client.SetBaseURL(baseUrl[1 : len(baseUrl)-1])

		var R Response
		err = chat.Chat.Client.Call("notification", "createNotification", notif, &R)
		if err != nil {
			log.Println(err)
		}
	}
}
