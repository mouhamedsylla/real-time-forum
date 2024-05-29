package server

import (
	"fmt"
	"log"
	"net/http"
	"real-time-forum/server/microservices"

	"github.com/mouhamedsylla/term-color/color"
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
	clr := color.Color()
	for _, service := range s.Services.Microservices {
		service := service.GetService()
		go func(svc *microservices.Microservice) {
			clr.SetText(fmt.Sprintf("[RUNNING SERVICE] %s", svc.ServiceName))
			clr.ColorTextPattern("[RUNNING SERVICE]", clr.Yellow)
			clr.ColorTextPattern(svc.ServiceName, clr.Green)
			fmt.Println(clr.ToString())
			server := http.Server{
				Addr:    svc.Port,
				Handler: svc.Router,
			}
			log.Fatalln(server.ListenAndServe())
		}(service)
	}
	select {}
}
