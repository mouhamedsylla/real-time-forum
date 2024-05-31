package auth

import (
	"net/http"
	"real-time-forum/utils"
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
	// data, err := io.ReadAll(rq.Body)
	// if err != nil {
	// 	http.Error(w, "500", http.StatusInternalServerError)
	// 	return
	// }

	// var userRegister userRegister

	// if err = json.Unmarshal(data, &userRegister); err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// if err = storage.Insert(userRegister); err != nil {
	// 	return
	// }

	data, status, err := utils.DecodeJSONRequestBody(rq, userRegister{})

	if err != nil {
		utils.ResponseWithJSON(w, err, status)
		return
	}

	user := data.(*userRegister)
	CryptPassword(user)
	if err = storage.Insert(*user); err != nil {
		utils.ResponseWithJSON(w, "Service Auth.Register: Bad Request", http.StatusBadRequest)
		return
	}

	utils.ResponseWithJSON(w, "Registering Successfuly", http.StatusOK)

}

// func validForm(s string) {

// }
