package auth

import (
	"encoding/json"
	"io"
	"net/http"
)

func (r *Register) HTTPServe() http.Handler {
	return http.HandlerFunc(r.Register)
}

func (r *Register) EndPoint() string {
	return "/register"
}

func (r *Register) SetMethods() []string {
	return []string{"POST"}
}

func (r *Register) Register(w http.ResponseWriter, rq *http.Request) {
	data, err := io.ReadAll(rq.Body)
	if err != nil {
		http.Error(w, "404", http.StatusBadRequest)
		return
	}

	var user_register userRegister

	if err = json.Unmarshal(data, &user_register); err != nil {
		return
	}
}

// func validForm(s string) {

// }
