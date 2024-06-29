package notification

import (
	"real-time-forum/server/microservices"
	"real-time-forum/server/middleware"
	"real-time-forum/services/notification/controllers"
	"real-time-forum/services/notification/database"
	"real-time-forum/services/notification/models"
	"real-time-forum/utils"
	"sync"
)

const (
	DB_NAME = "notification.db"
	DB_PATH = "../../services/notification/database/"
)

type Notification struct {
	Notification *microservices.Microservice
}

func (notif *Notification) ConfigureEndpoint() {
	for _, controller := range notif.Notification.Controllers {
		notif.Notification.Router.Method(controller.SetMethods()...).Middleware(
			middleware.LogRequest,
		).Handler(
			controller.EndPoint(),
			controller.HTTPServe(),
		)
	}
}

func (notif *Notification) InitService() (err error) {
	database.DbNotification.Storage, err = utils.InitStorage(DB_NAME, DB_PATH, models.UserNotification{})
	controller := []microservices.Controller{
		// add controller...
		&controllers.CreateNotification{},
		&controllers.SendNotification{},
		&controllers.ConnectedUser{},
	}

	notif.Notification = microservices.NewMicroservice("Notification", ":9191")
	notif.Notification.Controllers = append(notif.Notification.Controllers, controller...)
	go notif.HandleUserDisconnect()
	go notif.HandleUserConnect()
	return err
}

func (notif *Notification) GetService() *microservices.Microservice {
	return notif.Notification
}

func (notif *Notification) HandleUserDisconnect() {
	var mutex sync.Mutex
	for {
		mutex.Lock()
		idUser := <-controllers.Disconnect_channel
		mutex.Unlock()

		infos_user := models.UserInfos{
			Type:   "user_status",
			Id:     idUser,
			Status: "offline",
		}
		for _, client := range controllers.Clients {
			client.WriteJSON(infos_user)
		}
	}
}

func (notif *Notification) HandleUserConnect() {
	var mutex sync.Mutex
	for {
		mutex.Lock()
		idUser := <-controllers.Connection_channel
		mutex.Unlock()

		infos_user := models.UserInfos{
			Type:   "user_status",
			Id:     idUser,
			Status: "online",
		}
		for _, client := range controllers.Clients {
			client.WriteJSON(infos_user)
		}
	}
}
