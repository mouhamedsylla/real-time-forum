package microservices

import (
	"log"
	"real-time-forum/server/router"
)

type AppServices struct {
	Microservices []Service
}

type Microservice struct {
	ServiceName string
	Router      *router.Router
	Controllers []Controller
	Port        string
}

func (aps *AppServices) InitServices() {
	for _, service := range aps.Microservices {
		if err := service.InitService(); err != nil {
			log.Fatalf("Failed to initialize service %s: %v", service.GetService().ServiceName, err)
		}
		service.ConfigureEndpoint()
	}
}

func NewAppServices(services ...Service) *AppServices {
	aps := &AppServices{}
	aps.Microservices = append(aps.Microservices, services...)
	return aps
}

func NewMicroservice(name, port string) *Microservice {
	return &Microservice{
		ServiceName: name,
		Port:        port,
		Router:      router.NewRouter(),
	}
}
