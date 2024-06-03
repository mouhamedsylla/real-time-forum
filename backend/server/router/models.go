// Package router defines the structures and types used for HTTP routing.
package router

import "net/http"

// CustomsRoutes is a map defining custom routes with string keys and values.
type CustomsRoutes map[string]string

// Tree represents the routing tree with a root node.
type Tree struct {
	// Node is the root route of the routing tree.
	Node *Route
}

// Middleware is a function type that takes an http.Handler and returns an http.Handler.
// It is used for wrapping HTTP handlers with additional functionality.
type Middleware func(http.Handler) http.Handler

// Router represents the main router structure containing the routing tree,
// a temporary route for building routes, and static file serving configuration.
type Router struct {
	// Tree is the routing tree used to store and match routes.
	Tree *Tree
	
	// TempRoute is a temporary route used during route configuration.
	TempRoute Route
	
	// Static holds the configuration for serving static files.
	Static Directory
}

// Param represents a key-value pair used in dynamic routes.
type Param struct {
	Key   string // Key is the parameter name.
	Value string // Value is the parameter value.
}

// Route represents a single route in the routing tree.
type Route struct {
	Label      string              // Label is the name or description of the route.
	Methods    []string            // Methods is a list of HTTP methods (e.g., GET, POST) that the route handles.
	Handle     http.Handler        // Handle is the HTTP handler for the route.
	Child      map[string]*Route   // Child is a map of child routes.
	Middleware []Middleware        // Middleware is a list of middleware functions for the route.
	IsDynamic  bool                // IsDynamic indicates if the route is dynamic (contains parameters).
	Param      Param               // Param represents the key-value pair for the dynamic route.
}

// Directory represents the configuration for serving static files.
type Directory struct {
	Prefix string // Prefix is the URL prefix for accessing static files.
	Dir    http.Dir // Dir is the directory to serve static files from.
}