package auth

import (
	"net/http"
	"real-time-forum/utils"
)

func (r *Register) HTTPServe() http.Handler {
	return http.HandlerFunc(r.Register)
}

func (r *Register) EndPoint() string {
	return "/register/register"
}

func (r *Register) SetMethods() []string {
	return []string{"POST"}
}

func (r *Register) Register(w http.ResponseWriter, rq *http.Request) {
	data, status, err := utils.DecodeJSONRequestBody(rq, UserRegister{})

	if err != nil {
		utils.ResponseWithJSON(w, err, status)
		return
	}

	user := data.(*UserRegister)
	CryptPassword(user)
	if err = storage.Insert(*user); err != nil {
		utils.ResponseWithJSON(w, "Service Auth.Register: Bad Request", http.StatusBadRequest)
		return
	}

	utils.ResponseWithJSON(w, "Registering Successfuly", http.StatusOK)

}

// func validForm(s string) {

// }
