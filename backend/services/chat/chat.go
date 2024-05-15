package realtimechat

import (
	"net/http"
	"real-time-forum/server/microservices"
)

type Chat struct {
	Chat *microservices.Microservice
}

func (chat *Chat) ConfigureEndpoint() {
	for _, controller := range chat.Chat.Controllers {
		chat.Chat.Router.Method(http.MethodGet).Handler(controller.EndPoint(), controller.HTTPServe())
	}
}

func (chat *Chat) InitService() {
	controllers := []microservices.Controller{
		// add controller ...
	}
	chat.Chat = microservices.NewMicroservice("Realtime Chat", ":9090")
	chat.Chat.Controllers = append(chat.Chat.Controllers, controllers...)
}

func (chat *Chat) GetService() *microservices.Microservice {
	return chat.Chat
}
