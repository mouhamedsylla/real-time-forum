package controllers

import (
	"net/http"
	"real-time-forum/services/chat/database"
	"real-time-forum/services/chat/models"
	"real-time-forum/utils"
)


func (pmu *GetPrivateMessageUsers) HTTPServe() http.Handler {
	return http.HandlerFunc(pmu.getPrivateMessageUsers)
}

func (pmu *GetPrivateMessageUsers) EndPoint() string {
	return "/chat/message/private/users/:userId"
}

func (pmu *GetPrivateMessageUsers) SetMethods() []string {
	return []string{http.MethodGet}
}

func (pmu *GetPrivateMessageUsers) getPrivateMessageUsers(w http.ResponseWriter, r *http.Request) {
	CustomRoute := r.Context().Value("CustomRoute").(map[string]string)
	database.DbChat.Storage.Custom.Where("senderId", CustomRoute["userId"])
	result := database.DbChat.Storage.Scan(models.Message{}, "ReceiverId").([]models.Message)
	database.DbChat.Storage.Custom.Clear()
	utils.ResponseWithJSON(w, result, http.StatusOK)
}
