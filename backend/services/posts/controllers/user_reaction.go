package controllers

import (
	"net/http"
	"real-time-forum/services/posts/database"
	"real-time-forum/services/posts/models"
	"real-time-forum/utils"
	"strconv"
)

func (p *GetUserPostReactions) HTTPServe() http.Handler {
	return http.HandlerFunc(p.GetUserPostReactions)
}

func (p *GetUserPostReactions) EndPoint() string {
	return "/posts/GetUserPostReactions/:userId"
}

func (p *GetUserPostReactions) SetMethods() []string {
	return []string{"GET"}
}

func (p *GetUserPostReactions) GetUserPostReactions(w http.ResponseWriter, r *http.Request) {
	var response models.Response
	CustomRoute := r.Context().Value("CustomRoute").(map[string]string)
	userId, err := strconv.Atoi(CustomRoute["userId"])
	if err != nil {
		response.Message = "Id user not found"
		utils.ResponseWithJSON(w, response, http.StatusBadRequest)
		return
	}

	database.DbPost.Storage.Custom.Where("UserId", userId)
	data := database.DbPost.Storage.Scan(models.ReactionPost{}, "PostId", "UserId", "Value")
	database.DbPost.Storage.Custom.Clear()
	
	if data == nil {
		response.Message = "No reactions found"
		utils.ResponseWithJSON(w, response, http.StatusNotFound)
		return
	}

	result := data.([]models.ReactionPost)
	utils.ResponseWithJSON(w, result, http.StatusOK)
}
