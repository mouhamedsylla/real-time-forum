package chat

import (
	"net/http"
	"real-time-forum/utils"
)


func (pmu *getPrivateMessageUsers) HTTPServe() http.Handler {
	return http.HandlerFunc(pmu.getPrivateMessageUsers)
}

func (pmu *getPrivateMessageUsers) EndPoint() string {
	return "/chat/message/private/users/:userId"
}

func (pmu *getPrivateMessageUsers) SetMethods() []string {
	return []string{http.MethodGet}
}

func (pmu *getPrivateMessageUsers) getPrivateMessageUsers(w http.ResponseWriter, r *http.Request) {
	CustomRoute := r.Context().Value("CustomRoute").(map[string]string)
	storage.Custom.Where("senderId", CustomRoute["userId"])
	result := storage.Scan(Message{}, "ReceiverId").([]Message)
	storage.Custom.Clear()
	utils.ResponseWithJSON(w, result, http.StatusOK)
}
