package router

import (
	"errors"
	"net/http"
	"strings"
)

const (
	ROOT = "/"
)

func NewTree() *Tree {
	return &Tree{
		Node: &Route{
			Label: ROOT,
			Child: make(map[string]*Route),
		},
	}
}

func (t *Tree) Insert(path string, handler http.Handler, mid []Middleware, methods ...string) {
	actualRoute := t.Node
	buf := ""
	if path == ROOT {
		actualRoute.Methods = append(actualRoute.Methods, methods...)
		actualRoute.Middleware = mid
		actualRoute.Handle = handler
	} else {
		roads := strings.Split(path, "/")
		for _, routeName := range roads[1:] {
			if strings.HasPrefix(routeName, ":") {
				buf = routeName[1:]
				routeName = "?"
			}

			NextRoute, ok := actualRoute.Child[routeName]

			if !ok {
				actualRoute.Child[routeName] = NewRoute(routeName, mid, methods...)
				actualRoute = actualRoute.Child[routeName]
				if routeName == "?" {
					actualRoute.IsDynamic = true
					actualRoute.Param.Key = buf
				}
				continue
			}
			actualRoute = NextRoute
		}
		actualRoute.Handle = handler
	}
}

func (t *Tree) Search(method string, path string) (http.Handler, []Middleware, map[string]string, error) {
	actualRoute := t.Node
	custom_routes := make(map[string]string)
	if path != ROOT {
		roads := strings.Split(path, "/")
		for _, routeName := range roads[1:] {

			nextRoute, ok := actualRoute.Child[routeName]

			if !ok {
				nextRoute, ok = actualRoute.Child["?"]
				if !ok {
					break
				}
				actualRoute = nextRoute
				actualRoute.Param.Value = routeName
				custom_routes[actualRoute.Param.Key] = actualRoute.Param.Value
				continue
			}
			actualRoute = nextRoute
		}
	}

	if actualRoute.Handle == nil {
		err := errors.New(PAGE_NOT_FOUND)
		return nil, nil, nil, err
	}

	if err := actualRoute.IsAllowed(method); err != nil {
		return nil, nil, nil, err
	}

	return actualRoute.Handle, actualRoute.Middleware, custom_routes, nil
}
