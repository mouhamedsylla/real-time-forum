// Package router defines the structures and functions for managing a routing tree.
package router

import (
	"errors"
	"net/http"
	"strings"
)

// Constants
const (
	ROOT = "/"
)

// NewTree initializes and returns a new routing tree with a root route.
func NewTree() *Tree {
	return &Tree{
		Node: &Route{
			Label: ROOT,
			Child: make(map[string]*Route),
		},
	}
}

// Insert adds a new route to the routing tree with the specified path, handler, middleware, and methods.
func (t *Tree) Insert(path string, handler http.Handler, mid []Middleware, methods ...string) {
	actualRoute := t.Node
	buf := ""

	if path == ROOT {
		// If the path is the root, set the methods, middleware, and handler directly on the root route.
		actualRoute.Methods = append(actualRoute.Methods, methods...)
		actualRoute.Middleware = mid
		actualRoute.Handle = handler
	} else {
		// Split the path into segments and iterate over them to insert the route.
		roads := strings.Split(path, "/")
		for _, routeName := range roads[1:] {
			if strings.HasPrefix(routeName, ":") {
				// Handle dynamic route segments (e.g., ":id").
				buf = routeName[1:]
				routeName = "?"
			}

			NextRoute, ok := actualRoute.Child[routeName]

			if !ok {
				// If the route does not exist, create a new one.
				actualRoute.Child[routeName] = NewRoute(routeName, methods...)
				actualRoute = actualRoute.Child[routeName]
				if routeName == "?" {
					// Mark the route as dynamic and set the parameter key.
					actualRoute.IsDynamic = true
					actualRoute.Param.Key = buf
				}
				continue
			}
			// Move to the next route in the path.
			actualRoute = NextRoute
		}
		// Set the handler and middleware for the final route in the path.
		actualRoute.Handle = handler
		actualRoute.Middleware = mid
	}
}

// Search finds a route in the routing tree that matches the given method and path.
// Returns the handler, middleware, custom routes (dynamic parameters), and an error if any.
func (t *Tree) Search(method string, path string) (http.Handler, []Middleware, map[string]string, error) {
	actualRoute := t.Node
	custom_routes := make(map[string]string)

	if path != ROOT {
		// Split the path into segments and iterate over them to find the matching route.
		roads := strings.Split(path, "/")
		for _, routeName := range roads[1:] {
			nextRoute, ok := actualRoute.Child[routeName]

			if !ok {
				// Handle dynamic routes if the exact route is not found.
				nextRoute, ok = actualRoute.Child["?"]
				if !ok {
					break
				}
				actualRoute = nextRoute
				actualRoute.Param.Value = routeName
				custom_routes[actualRoute.Param.Key] = actualRoute.Param.Value
				continue
			}
			// Move to the next route in the path.
			actualRoute = nextRoute
		}
	}

	if actualRoute.Handle == nil {
		// If no handler is found for the route, return a PAGE_NOT_FOUND error.
		err := errors.New(PAGE_NOT_FOUND)
		return nil, nil, nil, err
	}

	// Check if the method is allowed for the found route.
	if err := actualRoute.IsAllowed(method); err != nil {
		return nil, nil, nil, err
	}

	return actualRoute.Handle, actualRoute.Middleware, custom_routes, nil
}
