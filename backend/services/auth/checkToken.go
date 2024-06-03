package auth

import (
	"net/http"
	"real-time-forum/utils"
)

func (r *CheckToken) HTTPServe() http.Handler {
	return http.HandlerFunc(r.CheckToken)
}

func (r *CheckToken) EndPoint() string {
	return "/auth/checkToken"
}

func (r *CheckToken) SetMethods() []string {
	return []string{"POST"}
}

func (r *CheckToken) CheckToken(w http.ResponseWriter, rq *http.Request) {
	m := Response{Message: "Token is valid"}
	data, status, err := utils.DecodeJSONRequestBody(rq, Request{})
	if err != nil {
		utils.ResponseWithJSON(w, err, status)
		return
	}
	token := data.(*Request)
	_, err = utils.VerifyToken(token.Token, jwt.Key.Public)

	if err != nil {
		m.Message = "Token is invalid"
		utils.ResponseWithJSON(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	utils.ResponseWithJSON(w, m, http.StatusOK)
}
