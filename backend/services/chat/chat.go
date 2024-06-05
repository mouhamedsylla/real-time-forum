package chat

import (
	"net/http"
	"real-time-forum/orm"
	"real-time-forum/server/microservices"
	"real-time-forum/server/middleware"
	"real-time-forum/utils"
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
			Middleware(middleware.LogRequest, middleware.Authenticate).
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
	return
}

func (chat *Chat) GetService() *microservices.Microservice {
	return chat.Chat
}
