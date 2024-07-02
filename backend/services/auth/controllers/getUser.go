package controllers

import (
	"net/http"
	"real-time-forum/services/auth/database"
	"real-time-forum/services/auth/models"
	"real-time-forum/utils"
	"strconv"
)

func (gtuser *GetUser) HTTPServe() http.Handler {
	return http.HandlerFunc(gtuser.GetUser)
}

func (gtuser *GetUser) EndPoint() string {
	return "/auth/getUsers"
}

func (gtuser *GetUser) SetMethods() []string {
	return []string{"GET"}
}

func (gtuser *GetUser) GetUser(w http.ResponseWriter, r *http.Request) {
	var result []models.UserRegister
	var response models.Response
	param := r.URL.Query().Get("userId")
	if param != "" {
		userId, _ := strconv.Atoi(param)
		database.Db.Storage.Custom.Where("Id", userId)
		data := database.Db.Storage.Scan(models.UserRegister{}, "Id", "Nickname", "FirstName", "LastName", "Email")
		database.Db.Storage.Custom.Clear()
		if data == nil {
			response.Message = "user not found"
			utils.ResponseWithJSON(w, response, http.StatusNotFound)
			return
		}
		result = data.([]models.UserRegister)
		user := result[0]
		utils.ResponseWithJSON(w, user, http.StatusOK)
		return
	}

	result = database.Db.Storage.Scan(models.UserRegister{}, "Id", "Nickname", "FirstName", "LastName", "Email").([]models.UserRegister)
	if len(result) == 0 {
		response.Message = "Any user found"
		utils.ResponseWithJSON(w, response, http.StatusNotFound)
		return
	}
	utils.ResponseWithJSON(w, result, http.StatusOK)
}
