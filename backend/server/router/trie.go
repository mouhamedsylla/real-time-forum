package router

import (
	"errors"
	"net/http"
	"strings"
)

const (
	ROOT = "/"
)

func NewParam(key string) Params {
	return Params{
		key: key,
	}
}

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
	if path == ROOT {
		actualRoute.Methods = append(actualRoute.Methods, methods...)
		actualRoute.Middleware = mid
		actualRoute.Handle = handler
	} else {
		roads := strings.Split(path, "/")
		for i, routeName := range roads[1:] {
			if strings.HasPrefix(routeName, ":") {
				actualRoute.IsDynamic = true
				actualRoute.Params = append(actualRoute.Params, NewParam(routeName[1:]))
				continue
			}
			NextRoute, ok := actualRoute.Child[routeName]
			if ok {
				actualRoute = NextRoute
			}
			if !ok {
				actualRoute.Child[routeName] = NewRoute(routeName, mid, methods...)
				actualRoute = actualRoute.Child[routeName]
			}

			if i == len(roads[1:])-1 {
				actualRoute.Handle = handler
			}
		}
	}
}

func (t *Tree) Search(method string, path string) (http.Handler, []Middleware, error) {
	actualRoute := t.Node
	k := 0
	if path != ROOT {
		roads := strings.Split(path, "/")
		for _, routeName := range roads[1:] {

			nextRoute, ok := actualRoute.Child[routeName]
			if !ok {
				if actualRoute.IsDynamic {
					if k >= len(actualRoute.Params) {
						err := errors.New(ROUTE_NOT_FOUND)
						return nil, nil, err
					}
					actualRoute.Params[k].value = routeName
					k++
					continue
				}
				if routeName == actualRoute.Label || actualRoute.IsDynamic {
					break
				} else {
					err := errors.New(ROUTE_NOT_FOUND)
					return nil, nil, err
				}
			}
			actualRoute = nextRoute

		}
	}
	if err := actualRoute.IsAllowed(method); err != nil {
		return nil, nil, err
	}

	if actualRoute.Handle == nil {
		err := errors.New(PAGE_NOT_FOUND)
		return nil, nil, err
	}
	return actualRoute.Handle, actualRoute.Middleware, nil
}
