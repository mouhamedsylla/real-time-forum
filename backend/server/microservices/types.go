package microservices

import (
	"net/http"
)

type Service interface {
	GetService() *Microservice
	InitService()
	ConfigureEndpoint()
}

type Controller interface {
	HTTPServe() http.Handler
	EndPoint() string
}
