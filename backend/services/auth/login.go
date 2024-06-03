package auth

import (
	"net/http"
	"real-time-forum/utils"
	"time"
)

func (l *Login) HTTPServe() http.Handler {
	return http.HandlerFunc(l.Login)
}

func (l *Login) EndPoint() string {
	return "/auth/login"
}

func (l *Login) SetMethods() []string {
	return []string{"POST"}
}

func (l *Login) Login(w http.ResponseWriter, r *http.Request) {
	data, status, err := utils.DecodeJSONRequestBody(r, userLogin{})
	if err != nil {
		utils.ResponseWithJSON(w, err, status)
		return
	}

	toAuthenticate := *data.(*userLogin)
	storage.Custom.Where("Email", &toAuthenticate.Identifier).Or("Nickname", &toAuthenticate.Identifier)
	rslt := storage.Scan(UserRegister{}, "Password").([]UserRegister)
	storage.Custom.Clear()

	if len(rslt) == 0 {
		utils.ResponseWithJSON(w, "this user does't exist", http.StatusNotFound)
		return
	}

	user := rslt[0]
	if err = Authenticate(user.Password, &toAuthenticate); err != nil {
		utils.ResponseWithJSON(w, err.Error(), http.StatusUnauthorized)
		return
	}

	token := jwt.GenerateToken()
	http.SetCookie(w, &http.Cookie{
		Name:    "forum",
		Value:   token,
		Expires: time.Now().Add(10 * time.Minute),
		Path:   "/",
	})

	m := Response{Message: "login successfull"}
	utils.ResponseWithJSON(w, m, http.StatusOK)
}
