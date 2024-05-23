package router

import "net/http"

type CustomsRoutes map[string]string
type Tree struct {
	Node *Route
}

type Middleware func(http.Handler) http.Handler

type Router struct {
	Tree      *Tree
	TempRoute Route
	Static    Directory
}

type Param struct {
	Key   string
	Value string
}
type Route struct {
	Label      string
	Methods    []string
	Handle     http.Handler
	Child      map[string]*Route
	Middleware []Middleware
	IsDynamic  bool
	Param      Param
}

type Directory struct {
	Prefix string
	Dir    http.Dir
}
