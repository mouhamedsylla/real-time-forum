package chat

import (
	"log"
	"net/http"
	"os"
	"real-time-forum/orm"
	"real-time-forum/server/microservices"
	"real-time-forum/server/middleware"
	"real-time-forum/utils"
	"strings"
)

type Request struct {
	Token string `json:"token"`
}

type Response struct {
	Message string `json:"message"`
}

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
		Middleware(middleware.LogRequest, chat.IsConnected).
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

func (chat *Chat) IsConnected(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		err := utils.LoadEnv("../../.env")
		if err != nil {
			log.Fatal(err.Error())
		}

		urlBase := os.Getenv("AUTH_SERVICE")
		chat.Chat.Client.SetBaseURL(strings.Replace(urlBase, `"`, "", -1))
		chat.Chat.Client.SetMethod(http.MethodPost)
		
		session, err := r.Cookie("forum")
		if err != nil {
			utils.ResponseWithJSON(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		request := Request{Token: session.Value}
		var R Response
		err = chat.Chat.Client.Call("auth", "checkToken", request, &R)
		if err != nil {
			utils.ResponseWithJSON(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
