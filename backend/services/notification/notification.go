package notification

import (
	"real-time-forum/server/microservices"
	"real-time-forum/server/middleware"
	"real-time-forum/services/notification/controllers"
	"real-time-forum/services/notification/database"
	"real-time-forum/services/notification/models"
	"real-time-forum/utils"
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
	}

	notif.Notification = microservices.NewMicroservice("Notification", ":9191")
	notif.Notification.Controllers = append(notif.Notification.Controllers, controller...)
	return err
}

func (notif *Notification) GetService() *microservices.Microservice {
	return notif.Notification
}
