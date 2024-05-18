package chat

import "net/http"

func (pmu *getPrivateMessageUsers) HTTPServe() http.Handler{
	return http.HandlerFunc(pmu.getPrivateMessageUsers)
}

func (pmu *getPrivateMessageUsers) EndPoint() string {
	return "/api/messages/private/users/:userId"
}

func (pmu *getPrivateMessageUsers) getPrivateMessageUsers(w http.ResponseWriter, r *http.Request) {

}