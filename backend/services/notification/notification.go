package notification

import (
	"real-time-forum/orm"
	"real-time-forum/server/microservices"
	"real-time-forum/server/middleware"
	"real-time-forum/utils"
)

const (
	DB_NAME = "notification.db"
	DB_PATH = "../../services/notification/database/"
)

var storage *orm.ORM

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
	storage, err = utils.InitStorage(DB_NAME, DB_PATH, UserNotification{})
	controller := []microservices.Controller{
		// add controller...
		&CreateNotification{},
	}

	notif.Notification = microservices.NewMicroservice("Notification", ":9191")
	notif.Notification.Controllers = append(notif.Notification.Controllers, controller...)
	return err
}

func (notif *Notification) GetService() *microservices.Microservice{
	return notif.Notification
}

