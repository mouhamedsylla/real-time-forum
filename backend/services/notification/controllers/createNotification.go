package controllers

import (
	"net/http"
	"real-time-forum/services/notification/database"
	"real-time-forum/services/notification/models"
	"real-time-forum/utils"
)

var userNotif = make(chan models.UserNotification)

func (cn *CreateNotification) HTTPServe() http.Handler {
	return http.HandlerFunc(cn.CreateNotification)
}

func (cn *CreateNotification) EndPoint() string {
	return "/notification/createNotification"
}

func (cn *CreateNotification) SetMethods() []string {
	return []string{"POST"}
}

func (cn *CreateNotification) CreateNotification(w http.ResponseWriter, rq *http.Request) {
	data, status, err := utils.DecodeJSONRequestBody(rq, models.UserNotification{})
	if err != nil {
		utils.ResponseWithJSON(w, err, status)
		return
	}
	notification := data.(*models.UserNotification)
	if err = database.DbNotification.Storage.Insert(*notification); err != nil {
		utils.ResponseWithJSON(w, "Service Notification.CreateNotification: Bad Request", http.StatusBadRequest)
		return
	}

	userNotif <- *notification
	message := models.Response{
		Message: "Notification Created",
	}

	utils.ResponseWithJSON(w, message, http.StatusOK)
}
