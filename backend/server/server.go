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
	var wg sync.WaitGroup
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	fmt.Println("🔁 Starting services...")

	for _, service := range s.Services.Microservices {
		service := service.GetService()
		wg.Add(1)
		go func(svc *microservices.Microservice) {
			defer wg.Done()
			fmt.Printf("[RUNNING SERVICE] ✅ %s\n", svc.ServiceName)

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
			
			fmt.Printf("[SHUTTING DOWN SERVICE] 🛑 %s\n", svc.ServiceName)
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			if err := server.Shutdown(ctx); err != nil {
				log.Fatalf("Could not gracefully shutdown the server %s: %v\n", svc.ServiceName, err)
			}
		}(service)
	}
	wg.Wait()
}
