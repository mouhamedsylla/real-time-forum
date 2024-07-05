// Package microservices defines interfaces for building microservice components
package microservices

import (
	"net/http"
)

// Service represents the interface for a microservice.
// It includes methods to retrieve the service instance, initialize the service,
// and configure its endpoints.
type Service interface {
	// GetService returns the Microservice instance.
	GetService() *Microservice

	// InitService initializes the service and returns an error if initialization fails.
	InitService() error

	// ConfigureEndpoint sets up the service endpoints.
	ConfigureEndpoint()
}

// Controller represents the interface for a microservice controller.
// It includes methods to serve HTTP requests, define allowed HTTP methods,
// and specify the endpoint path.
type Controller interface {
	// HTTPServe returns an http.Handler to handle HTTP requests.
	HTTPServe() http.Handler

	// SetMethods defines and returns a slice of HTTP methods (e.g., GET, POST) that the controller supports.
	SetMethods() []string

	// EndPoint returns the URL endpoint as a string.
	EndPoint() string
}

// Client represents the interface for a microservice client.
// It includes methods to call other services, set the base URL for the client,
// and specify the HTTP method for requests.
type Client interface {
	// Call makes a request to another service.
	// It takes the service name, endpoint, request data, and a response structure, returning an error if the call fails.
	Call(serviceName, endpoint string, request, response interface{}) error

	// SetBaseURL sets the base URL for the client.
	SetBaseURL(url string)

	// SetMethod sets the HTTP method (e.g., GET, POST) for the client.
	SetMethod(method string)
}
