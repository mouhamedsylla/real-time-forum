package router

import "net/http"

type Tree struct {
	Node *Route
}

type Middleware func(http.Handler) http.Handler

type Router struct {
	Tree         *Tree
	TempRoute Route
	Static Directory
}

type Params struct {
	key string
	value string
}
type Route struct {
	Label   string
	Methods []string
	Handle  http.Handler
	Child   map[string]*Route
	Middleware []Middleware
	IsDynamic bool
	Params []Params
}

type Directory struct {
	Prefix string
	Dir http.Dir
}
