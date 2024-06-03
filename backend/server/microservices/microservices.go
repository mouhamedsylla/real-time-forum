// Package microservices defines the structures and functions for initializing and managing microservices.
package microservices

import (
	"log"
	"real-time-forum/server/router"
)

// AppServices is a container for multiple microservices.
type AppServices struct {
	// Microservices holds a list of Service interfaces representing different microservices.
	Microservices []Service
}

// Microservice defines the structure of an individual microservice.
type Microservice struct {
	// ServiceName is the name of the microservice.
	ServiceName string
	
	// Router handles the HTTP routing for the microservice.
	Router *router.Router
	
	// Controllers is a list of Controller interfaces for handling specific routes.
	Controllers []Controller
	
	// Port is the port number on which the microservice will run.
	Port string
	
	// Client is used for making requests to other services.
	Client Client
}

// InitServices initializes all microservices in AppServices.
// It iterates over each service, initializes it, and configures its endpoints.
func (aps *AppServices) InitServices() {
	for _, service := range aps.Microservices {
		// Initialize the service and handle any errors.
		if err := service.InitService(); err != nil {
			log.Fatalf("Failed to initialize service %s: %v", service.GetService().ServiceName, err)
		}
		// Configure the endpoints for the service.
		service.ConfigureEndpoint()
	}
}

// NewAppServices creates a new AppServices instance and adds the provided services to it.
// This function takes a variable number of Service interfaces as arguments.
func NewAppServices(services ...Service) *AppServices {
	aps := &AppServices{}
	aps.Microservices = append(aps.Microservices, services...)
	return aps
}

// NewMicroservice creates a new Microservice instance with the specified name and port.
// It initializes the Router and Client for the microservice.
func NewMicroservice(name, port string) *Microservice {
	return &Microservice{
		ServiceName: name,
		Port:        port,
		Router:      router.NewRouter(),
		Client:      NewClient(),
	}
}

