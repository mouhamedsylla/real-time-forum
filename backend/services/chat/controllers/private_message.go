package controllers

import (
	"net/http"
	"real-time-forum/services/chat/database"
	"real-time-forum/services/chat/models"
	"real-time-forum/utils"
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
	database.DbChat.Storage.Custom.
		Where("senderId", CustomRoute["senderId"]).
		And("receiverId", CustomRoute["receiverId"])

	result := database.DbChat.Storage.Scan(models.Message{}, "SenderId", "ReceiverId", "Content").([]models.Message)
	database.DbChat.Storage.Custom.Clear()
	utils.ResponseWithJSON(w, result, http.StatusOK)
}
