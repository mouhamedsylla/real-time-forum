package chat

import (
	"fmt"
	"net/http"
)

func (pm *getPrivateMessage) HTTPServe() http.Handler {
	return http.HandlerFunc(pm.getPrivateMessage)
}

func (pm *getPrivateMessage) EndPoint() string {
	return "/api/message/private/:senderId/:receiverId"
}

func (pm *getPrivateMessage) getPrivateMessage(w http.ResponseWriter, r *http.Request) {
	CustomRoute := r.Context().Value("CustomRoute")
	fmt.Println(CustomRoute)
}
