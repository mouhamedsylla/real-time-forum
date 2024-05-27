package auth

import (
	"net/http"
	"real-time-forum/server/microservices"
	"real-time-forum/utils"
)

const (
	DB_NAME = "auth.db"
	DB_PATH = "../services/auth/database/"
)

var stockage *orm.ORM

type Auth struct {
	Auth *microservices.Microservice
}

func (auth *Auth) ConfigureEndpoint() {
	for _, controller := range auth.Auth.Controllers {
		auth.Auth.Router.Method(controller.SetMethods).Handler(
			controller.EndPoint(), 
			controller.HTTPServe(),
		)
	}
}

func (auth *Auth) InitService() {
	stockage = utils.InitStorage(DB_NAME, DB_PATH, userRegister{}, userLogin{})
	controllers := []microservices.Controller{
		// add controller ...
		&Register{},
	}
	auth.Auth = microservices.NewMicroservice("Authentication", ":8080")
	auth.Auth.Controllers = append(auth.Auth.Controllers, controllers...)
}

func (auth *Auth) GetService() *microservices.Microservice {
	return auth.Auth
}