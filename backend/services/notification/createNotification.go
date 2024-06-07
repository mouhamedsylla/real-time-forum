package notification

import (
	"net/http"
	"real-time-forum/utils"
)

var userNotif = make(chan UserNotification)

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
	data, status, err := utils.DecodeJSONRequestBody(rq, UserNotification{})
	if err != nil {
		utils.ResponseWithJSON(w, err, status)
		return
	}
	notification := data.(*UserNotification)
	if err = storage.Insert(*notification); err != nil {
		utils.ResponseWithJSON(w, "Service Notification.CreateNotification: Bad Request", http.StatusBadRequest)
		return
	}

	userNotif <- *notification
	message := Response{
		Message: "Notification Created",
	}

	utils.ResponseWithJSON(w, message, http.StatusOK)
}
