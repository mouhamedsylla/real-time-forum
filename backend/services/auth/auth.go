package auth

import (
	"net/http"
	"real-time-forum/server/microservices"
)

type Auth struct {
	Auth *microservices.Microservice
}

func (auth *Auth) ConfigureEndpoint() {
	for _, controller := range auth.Auth.Controllers {
		auth.Auth.Router.Method(http.MethodGet).Handler(controller.EndPoint(), controller.HTTPServe())
	}
}

func (auth *Auth) InitService() {
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