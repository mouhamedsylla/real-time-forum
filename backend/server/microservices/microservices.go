package microservices

import (
	"real-time-forum/server/router"
)

type AppServices struct {
	Microservices []Service
}

func (aps *AppServices) InitServices() {
	for _, service := range aps.Microservices {
		service.InitService()
		service.ConfigureEndpoint()
	}
}

func NewAppServices(services ...Service) *AppServices {
	aps := &AppServices{}
	aps.Microservices = append(aps.Microservices, services...)
	return aps
}

type Microservice struct {
	ServiceName string
	Router      *router.Router
	Controllers []Controller
	Port        string
}

func NewMicroservice(name, port string) *Microservice {
	return &Microservice{
		ServiceName: name,
		Port:        port,
		Router:      router.NewRouter(),
	}
}
