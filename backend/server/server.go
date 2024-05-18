package server

import (
	"log"
	"net/http"
	"real-time-forum/server/microservices"
)

type Server struct {
	Services *microservices.AppServices
}

func NewServer(services ...microservices.Service) *Server {
	return &Server{
		Services: microservices.NewAppServices(services...),
	}
}

func (s *Server) StartServices() {
	s.Services.InitServices()
	for _, service := range s.Services.Microservices {
		service := service.GetService()
		go func(svc *microservices.Microservice) {
			log.Printf("%v service starting in: http://localhost%v", service.ServiceName, service.Port)
			server := http.Server{
				Addr:    svc.Port,
				Handler: svc.Router,
			}
			log.Fatalln(server.ListenAndServe())
		}(service)
	}
	select {}
}
