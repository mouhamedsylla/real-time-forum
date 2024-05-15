package router

import (
	"errors"
	"net/http"
	"strings"
)

const (
	METHOD_NOT_ALLOWED = "method not allowed"
	ROUTE_NOT_FOUND    = "route not found"
	PAGE_NOT_FOUND = "page not found"
)

func NewRouter() *Router {
	return &Router{
		Tree:         NewTree(),
		TempRoute: Route{},
	}
}

func NewRoute(label string, mid []Middleware, methods ...string) *Route {
	return &Route{
		Label:      label,
		Methods:    methods,
		Child:      make(map[string]*Route),
		Middleware: mid,
	}
}

func (R *Router) Method(methods ...string) *Router {
	R.TempRoute.Methods = methods
	return R
}

func (R *Router) Handler(path string, handler http.Handler) {
	R.TempRoute.Handle = handler
	R.Tree.Insert(path, R.TempRoute.Handle, R.TempRoute.Middleware, R.TempRoute.Methods...)
	R.TempRoute.Middleware = []Middleware{}
}

func (R *Router) Middleware(m ...Middleware) *Router {
	R.TempRoute.Middleware = m
	return R
}

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

func (R *Router) SetDirectory(prefix string, dir string) {
	R.Static.Prefix = prefix
	R.Static.Dir = http.Dir(dir)
}

func (R *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	path := r.URL.Path
	if strings.Contains(path, R.Static.Prefix) && R.Static.Prefix != ""{
		path = R.Static.Prefix
	}
	handler, middlewares, err := R.Tree.Search(method, path)
	if err != nil {
		status, msg := HandleError(err)
		w.WriteHeader(status)
		w.Write([]byte(msg))
		return
	}

	if len(middlewares) > 0 {
		for _, middleware := range middlewares {
			handler = middleware(handler)
		}
	}
	handler.ServeHTTP(w, r)
}

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

func (r *Route) IsAllowed(method string) error {
	for _, m := range r.Methods {
		if m == method {
			return nil
		}
	}
	return errors.New(METHOD_NOT_ALLOWED)
}
