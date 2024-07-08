package controllers

import (
	"net/http"
	"os"
	"real-time-forum/services/auth/database"
	"real-time-forum/services/auth/models"
	"real-time-forum/utils"
	"real-time-forum/utils/jwt"
	"time"
)

var Jwt = jwt.JWT{}

func (l *Login) HTTPServe() http.Handler {
	return http.HandlerFunc(l.Login)
}

func (l *Login) EndPoint() string {
	return "/auth/public/login"
}

func (l *Login) SetMethods() []string {
	return []string{"POST"}
}

func (l *Login) Login(w http.ResponseWriter, r *http.Request) {
	var response models.Response
	data, status, err := utils.DecodeJSONRequestBody(r, models.UserLogin{})
	if err != nil {
		response.Message = err.Error()
		utils.ResponseWithJSON(w, response, status)
		return
	}

	toAuthenticate := *data.(*models.UserLogin)
	database.Db.Storage.Custom.Where("Email", &toAuthenticate.Identifier).Or("Nickname", &toAuthenticate.Identifier)
	rslt := database.Db.Storage.Scan(models.UserRegister{}, "Password", "Id", "Nickname", "FirstName", "LastName", "Email").([]models.UserRegister)
	database.Db.Storage.Custom.Clear()

	if len(rslt) == 0 {
		response.Message = "this user does't exist"
		utils.ResponseWithJSON(w, response, http.StatusNotFound)
		return
	}

	user := rslt[0]
	if err = models.Authenticate(user.Password, &toAuthenticate); err != nil {
		response.Message = err.Error()
		utils.ResponseWithJSON(w, response, http.StatusUnauthorized)
		return
	}
	var token string
	if err = GetUserToken(&token, user.Id); err != nil {
		response.Message = err.Error()
		utils.ResponseWithJSON(w, response, http.StatusInternalServerError)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:    "forum",
		Value:   token,
		Expires: time.Now().Add(60 * time.Minute),
		Path:    "/",
	})

	response.Message = "login successfull"
	utils.ResponseWithJSON(w, response, http.StatusOK)
}

func GetUserToken(token *string, userId int) error {
	key := jwt.Key{}
	if _, err := os.Stat("../../utils/key/private_key.pem"); os.IsNotExist(err) {
		if err := key.GenerateKey(); err != nil {
			return err
		}
		pemK := key.PEMfromKey()
		if err = pemK.SetPEMToFile("../../utils/key"); err != nil {
			return err
		}
	}
	if err := key.KeyfromPrivateFile("../../utils/key/private_key.pem"); err != nil {
		return err
	}

	*token = Jwt.GenerateToken(userId, key.Private)
	return nil
}
