package controllers

import (
	"net/http"
	"real-time-forum/services/notification/models"
	"real-time-forum/utils"
	"strconv"
)

func (pmu *ConnectedUser) HTTPServe() http.Handler {
	return http.HandlerFunc(pmu.CoonectedUser)
}

func (pmu *ConnectedUser) EndPoint() string {
	return "/chat/message/private/getConnectedUser/:userId"
}

func (pmu *ConnectedUser) SetMethods() []string {
	return []string{http.MethodGet}
}

func (pmu *ConnectedUser) CoonectedUser(w http.ResponseWriter, r *http.Request) {
	CustomRoute := r.Context().Value("CustomRoute").(map[string]string)
	userID, _ := strconv.Atoi(CustomRoute["userId"])
	var response []models.ConnectedUser
	for id := range Clients {
		if id != userID {
			response = append(response, models.ConnectedUser{Id: id})
		}
	}

	utils.ResponseWithJSON(w, response, http.StatusOK)
}
