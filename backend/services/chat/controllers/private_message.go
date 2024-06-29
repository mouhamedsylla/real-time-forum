package controllers

import (
	"net/http"
	"real-time-forum/services/chat/database"
	"real-time-forum/services/chat/models"
	"real-time-forum/utils"
	"strconv"
)

func (pm *GetPrivateMessage) HTTPServe() http.Handler {
	return http.HandlerFunc(pm.getPrivateMessage)
}

func (pm *GetPrivateMessage) EndPoint() string {
	return "/chat/message/private/:senderId/:receiverId"
}

func (pm *GetPrivateMessage) SetMethods() []string {
	return []string{http.MethodGet}
}

func (pm *GetPrivateMessage) getPrivateMessage(w http.ResponseWriter, r *http.Request) {
	CustomRoute := r.Context().Value("CustomRoute").(map[string]string)

	sendId, _ := strconv.Atoi(CustomRoute["senderId"])
	receiveId, _ := strconv.Atoi(CustomRoute["receiverId"])
	database.DbChat.Storage.Custom.
		Where("senderId", sendId).And("receiverId", receiveId).
		Or("senderId", receiveId).And("receiverId", sendId)

	result := database.DbChat.Storage.Scan(models.Message{},"Id", "CreatedAt", "SenderId", "ReceiverId", "Content").([]models.Message)
	database.DbChat.Storage.Custom.Clear()
	utils.ResponseWithJSON(w, result, http.StatusOK)
}
