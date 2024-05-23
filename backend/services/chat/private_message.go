package chat

import (
	"net/http"
	"real-time-forum/utils"
)

func (pm *getPrivateMessage) HTTPServe() http.Handler {
	return http.HandlerFunc(pm.getPrivateMessage)
}

func (pm *getPrivateMessage) EndPoint() string {
	return "/api/message/private/:senderId/:receiverId"
}

func (pm *getPrivateMessage) getPrivateMessage(w http.ResponseWriter, r *http.Request) {
	CustomRoute := r.Context().Value("CustomRoute").(map[string]string)
	storage.Custom.
		Where("senderId", CustomRoute["senderId"]).
		And("receiverId", CustomRoute["receiverId"])

	result := storage.Scan(Message{}, "SenderId", "ReceiverId", "Content").([]Message)
	storage.Custom.Clear()
	utils.RespondWithJSON(w, result, http.StatusOK)
}
