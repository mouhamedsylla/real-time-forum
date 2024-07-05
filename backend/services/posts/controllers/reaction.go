package controllers

import (
	"net/http"
	"real-time-forum/services/posts/database"
	"real-time-forum/services/posts/models"
	"real-time-forum/utils"
	"strconv"
)

func (p *ReactionPosts) HTTPServe() http.Handler {
	return http.HandlerFunc(p.ReactionPosts)
}

func (p *ReactionPosts) EndPoint() string {
	return "/posts/ReactionPosts/:userId/:postId"
}

func (p *ReactionPosts) SetMethods() []string {
	return []string{"POST"}
}

func (p *ReactionPosts) ReactionPosts(w http.ResponseWriter, r *http.Request) {
	var response models.Response
	CustomRoute := r.Context().Value("CustomRoute").(map[string]string)
	postId, err := strconv.Atoi(CustomRoute["postId"])

	if err != nil {
		response.Message = "Id post not found"
		utils.ResponseWithJSON(w, response, http.StatusBadRequest)
		return
	}

	userId, err := strconv.Atoi(CustomRoute["userId"])
	if err != nil {
		response.Message = "Id user not found"
		utils.ResponseWithJSON(w, response, http.StatusBadRequest)
		return
	}

	data, status, err := utils.DecodeJSONRequestBody(r, models.ReactionPost{})
	if err != nil {
		utils.ResponseWithJSON(w, err, status)
		return
	}

	value := data.(*models.ReactionPost)

	err = database.UpdateReaction(database.DbPost.Storage.Db, userId, postId, value.Value)

	if err != nil {
		response.Message = "Error updating reaction"
		utils.ResponseWithJSON(w, response, http.StatusBadRequest)
		return
	}

	response.Message = "Reaction updated"
	utils.ResponseWithJSON(w, response, http.StatusOK)
}
