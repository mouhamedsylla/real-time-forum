package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"real-time-forum/server/microservices"
	"sync"
	"syscall"
	"time"

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
	var wg sync.WaitGroup
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	for _, service := range s.Services.Microservices {
		service := service.GetService()
		wg.Add(1)
		go func(svc *microservices.Microservice) {
			defer wg.Done()
			clr.SetText(fmt.Sprintf("[RUNNING SERVICE] ✅ %s", svc.ServiceName))
			clr.ColorTextPattern("[RUNNING SERVICE]", clr.Yellow)
			clr.ColorTextPattern(svc.ServiceName, clr.Green)
			fmt.Println(clr.ToString())

			server := http.Server{
				Addr:    svc.Port,
				Handler: svc.Router,
			}

			go func() {
				if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
					log.Fatalf("Could not listen on %s: %v\n", svc.Port, err)
				}
			}()
		
			<-stop
			
			fmt.Println(clr.SetText(fmt.Sprintf("[SHUTTING DOWN SERVICE] 🛑 %s", svc.ServiceName)).Colorize(clr.Red))
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			if err := server.Shutdown(ctx); err != nil {
				log.Fatalf("Could not gracefully shutdown the server %s: %v\n", svc.ServiceName, err)
			}
		}(service)
	}
	wg.Wait()
}
