package auth

import (
	"real-time-forum/server/microservices"
	"real-time-forum/services/auth/controllers"
	"real-time-forum/services/auth/database"
	"real-time-forum/services/auth/models"
	"real-time-forum/utils"
)

const (
	DB_NAME = "auth.db"
	DB_PATH = "../../services/auth/database/"
)

type Auth struct {
	Auth *microservices.Microservice
}

func (auth *Auth) ConfigureEndpoint() {
	for _, controller := range auth.Auth.Controllers {
		auth.Auth.Router.Method(controller.SetMethods()...).Handler(
			controller.EndPoint(),
			controller.HTTPServe(),
		)
	}
}

func (auth *Auth) InitService() (err error) {
	database.Db.Storage, err = utils.InitStorage(DB_NAME, DB_PATH, models.UserLogin{}, models.UserRegister{})
	controller := []microservices.Controller{
		// add controller ...
		&controllers.Register{},
		&controllers.Login{},
		&controllers.GetUser{},
		&controllers.GetGroupUserDiscussion{},
	}
	auth.Auth = microservices.NewMicroservice("Authentication", ":8080")
	controllers.AuthClient = auth.Auth.Client
	auth.Auth.Controllers = append(auth.Auth.Controllers, controller...)

	return
}

func (auth *Auth) GetService() *microservices.Microservice {
	return auth.Auth
}
