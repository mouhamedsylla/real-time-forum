package controllers

import (
	"net/http"
	"real-time-forum/services/chat/database"
	"real-time-forum/services/chat/models"
	"real-time-forum/utils"
	"strconv"
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
	senderId, _ := strconv.Atoi(CustomRoute["userId"])
	database.DbChat.Storage.Custom.Where("senderId", senderId)
	result := database.DbChat.Storage.Scan(models.Message{}, "ReceiverId").([]models.Message)
	database.DbChat.Storage.Custom.Clear()


	var rsp models.UserContact
	for _, v := range result {
		if ContainsInt(rsp.UsersId, v.ReceiverId) {
			continue
		}
		rsp.UsersId = append(rsp.UsersId, v.ReceiverId)
	}

	utils.ResponseWithJSON(w, rsp, http.StatusOK)
}

func ContainsInt(array []int, target int) bool {
	for _, value := range array {
		if value == target {
			return true
		}
	}
	return false
}
