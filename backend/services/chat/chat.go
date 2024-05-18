package chat

import (
	"net/http"
	"os"
	"real-time-forum/orm"
	"real-time-forum/server/microservices"
	"real-time-forum/utils"
)

const (
	DB_NAME = "chat.db"
	DB_PATH = "../services/chat/database/"
)

type Chat struct {
	Chat    *microservices.Microservice
	Storage *orm.ORM
}

func (chat *Chat) InitStorage() {
	if _, err := os.Stat(DB_PATH + DB_NAME); os.IsNotExist(err) {
		utils.CreateDatabase(DB_NAME, DB_PATH, Message{})
	}
	chat.Storage = utils.OrmInit(DB_NAME, DB_PATH)
}

func (chat *Chat) ConfigureEndpoint() {
	for _, controller := range chat.Chat.Controllers {
		chat.Chat.Router.Method(http.MethodGet).Handler(controller.EndPoint(), controller.HTTPServe())
	}
}

func (chat *Chat) InitService() {
	chat.InitStorage()
	controllers := []microservices.Controller{
		// add controller ...
		&sendMessage{},
		&getPrivateMessage{},
		&getPrivateMessageUsers{},
	}

	chat.Chat = microservices.NewMicroservice("Realtime Chat", ":9090")
	chat.Chat.Controllers = append(chat.Chat.Controllers, controllers...)
}

func (chat *Chat) GetService() *microservices.Microservice {
	return chat.Chat
}
