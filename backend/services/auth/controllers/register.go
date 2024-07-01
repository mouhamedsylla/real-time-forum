package controllers

import (
	"net/http"
	"real-time-forum/services/auth/database"
	"real-time-forum/services/auth/models"
	"real-time-forum/utils"
	validation "real-time-forum/utils/Validation"
)

func (r *Register) HTTPServe() http.Handler {
	return http.HandlerFunc(r.Register)
}

func (r *Register) EndPoint() string {
	return "/auth/public/register"
}

func (r *Register) SetMethods() []string {
	return []string{"POST"}
}

func (r *Register) Register(w http.ResponseWriter, rq *http.Request) {
	var response models.Response
	data, status, err := utils.DecodeJSONRequestBody(rq, models.UserRegister{})

	if err != nil {
		response.Message = err.Error()
		utils.ResponseWithJSON(w, response, status)
		return
	}

	user := data.(*models.UserRegister)
	var valid = validation.NewValidator()
	valid.Init(*user)

	if err := valid.Validate(); err != nil {
		response.Message = err.Error()
		utils.ResponseWithJSON(w, response, http.StatusBadRequest)
		return
	}

	models.CryptPassword(user)

	if err = database.Db.Storage.Insert(*user); err != nil {
		response.Message = err.Error()
		utils.ResponseWithJSON(w, response, http.StatusBadRequest)
		return
	}

	response.Message = "Registering Successfuly"
	utils.ResponseWithJSON(w, response, http.StatusOK)

}
