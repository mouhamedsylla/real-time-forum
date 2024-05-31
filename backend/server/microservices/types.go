package microservices

import (
	"net/http"
)

// The type `Service` defines methods for getting a microservice, initializing a service, and
// configuring an endpoint.
// @property GetService - The `GetService` method in the `Service` interface is expected to return a
// pointer to a `Microservice` object.
// @property InitService - The `InitService` method is used to initialize the service and set up any
// necessary configurations or dependencies before the service is started or used. This method
// typically includes tasks such as setting up database connections, initializing logging, loading
// configuration settings, and any other setup tasks required for the service to function properly.
// @property ConfigureEndpoint - The `ConfigureEndpoint` method is likely used to set up and configure
// the endpoint for the service. This could involve defining the route, handling different HTTP
// methods, setting up middleware, and any other necessary configurations for the endpoint to function
// correctly.
type Service interface {
	GetService() *Microservice
	InitService() error
	ConfigureEndpoint()
}

// The above type defines an interface for controllers in Go that includes methods for serving HTTP
// requests and returning the endpoint.
// @property HTTPServe - The `HTTPServe` method should return an `http.Handler` that will handle
// incoming HTTP requests for the controller.
// @property {string} EndPoint - The `EndPoint` property in the `Controller` interface represents the
// endpoint or URL path at which the controller will be accessible in the web application.
type Controller interface {
	HTTPServe() http.Handler
	SetMethods() []string
	EndPoint() string
}
