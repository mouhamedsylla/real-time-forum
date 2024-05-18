package auth

import (
	"fmt"
	"net/http"
)

func (r *Register) HTTPServe() http.Handler {
	return http.HandlerFunc(r.Register)
}

func (r *Register) EndPoint() string {
	return "/register"
}

func (r *Register) Register(w http.ResponseWriter, rq *http.Request) {
	fmt.Fprintf(w, "hello world")
}
