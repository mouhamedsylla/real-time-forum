// Package router defines the structures and functions for managing HTTP routes.
package router

import (
	"context"
	"errors"
	"net/http"
	"real-time-forum/server/middleware"
	"strings"
)

// Constants for error messages
const (
	METHOD_NOT_ALLOWED = "method not allowed"
	ROUTE_NOT_FOUND    = "route not found"
	PAGE_NOT_FOUND     = "page not found"
)

// NewRouter creates a new Router instance with an initialized routing tree and temporary route.
func NewRouter() *Router {
	return &Router{
		Tree:      NewTree(),
		TempRoute: Route{},
	}
}

// NewRoute creates a new Route instance with the specified label and HTTP methods.
func NewRoute(label string, methods ...string) *Route {
	return &Route{
		Label:   label,
		Methods: methods,
		Child:   make(map[string]*Route),
		Param:   Param{},
	}
}

// Method sets the HTTP methods for the temporary route in the Router and returns the Router.
func (R *Router) Method(methods ...string) *Router {
	R.TempRoute.Methods = methods
	return R
}

// Handler sets the handler for the specified path, inserts the route into the routing tree,
// and clears the temporary route's middleware.
func (R *Router) Handler(path string, handler http.Handler) {
	R.TempRoute.Handle = handler
	R.Tree.Insert(path, R.TempRoute.Handle, R.TempRoute.Middleware, R.TempRoute.Methods...)
	R.TempRoute.Middleware = []Middleware{}
}

// Middleware sets the middleware for the temporary route in the Router and returns the Router.
func (R *Router) Middleware(m ...Middleware) *Router {
	R.TempRoute.Middleware = m
	return R
}

// StaticServe returns an http.Handler for serving static files based on the Router's static file configuration.
func (R *Router) StaticServe() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		filePath := r.URL.Path[len(R.Static.Prefix):]
		file, err := R.Static.Dir.Open(filePath)
		if err != nil {
			http.NotFound(w, r)
			return
		}
		defer file.Close()

		fileInfo, err := file.Stat()
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		http.ServeContent(w, r, filePath, fileInfo.ModTime(), file)
	})
}

// SetDirectory sets the directory configuration for serving static files in the Router.
func (R *Router) SetDirectory(prefix string, dir string) {
	R.Static.Prefix = prefix
	R.Static.Dir = http.Dir(dir)
}

// ServeHTTP handles incoming HTTP requests, applying CORS middleware and routing the request
// to the appropriate handler based on the routing tree.
func (R *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	corsHandler := middleware.CORSMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		method := r.Method
		path := r.URL.Path

		if strings.Contains(path, R.Static.Prefix) && R.Static.Prefix != "" {
			path = R.Static.Prefix
		}

		handler, middlewares, custom_routes, err := R.Tree.Search(method, path)
		if err != nil {
			status, msg := HandleError(err)
			w.WriteHeader(status)
			w.Write([]byte(msg))
			return
		}

		for _, middleware := range middlewares {
			handler = middleware(handler)
		}

		ctx := context.WithValue(r.Context(), "CustomRoute", custom_routes)
		handler.ServeHTTP(w, r.WithContext(ctx))
	}))

	corsHandler.ServeHTTP(w, r)
}

// HandleError maps an error to an HTTP status code and message.
func HandleError(err error) (status int, msg string) {
	switch err.Error() {
	case METHOD_NOT_ALLOWED:
		status = http.StatusMethodNotAllowed
		msg = METHOD_NOT_ALLOWED
	case ROUTE_NOT_FOUND:
		status = http.StatusNotFound
		msg = ROUTE_NOT_FOUND
	case PAGE_NOT_FOUND:
		status = http.StatusNotFound
		msg = PAGE_NOT_FOUND
	}
	return
}

// IsAllowed checks if the given HTTP method is allowed for the route.
func (r *Route) IsAllowed(method string) error {
	for _, m := range r.Methods {
		if m == method {
			return nil
		}
	}
	return errors.New(METHOD_NOT_ALLOWED)
}
