package controllers

import (
	"net/http"
	"real-time-forum/services/posts/database"
	"real-time-forum/services/posts/models"
	"real-time-forum/utils"
)

func (p *CreatedPost) HTTPServe() http.Handler {
	return http.HandlerFunc(p.CreatedPost)
}

func (p *CreatedPost) EndPoint() string {
	return "/posts/createdpost"
}

func (p *CreatedPost) SetMethods() []string {
	return []string{"POST"}
}

func (p *CreatedPost) CreatedPost(w http.ResponseWriter, r *http.Request) {
	// Decode the JSON request body
	data, status, err := utils.DecodeJSONRequestBody(r, models.UserPosts{})
	if err != nil {
		response := models.ErrorResponse{Error: err.Error()}
		utils.ResponseWithJSON(w, response, status)
		return
	}

	// Type assert the decoded data to *UserPosts
	post, ok := data.(*models.UserPosts)
	if !ok {
		response := models.ErrorResponse{Error: "Invalid data format"}
		utils.ResponseWithJSON(w, response, http.StatusBadRequest)
		return
	}

	// Insert the post into the storage
	if err = database.DbPost.Storage.Insert(*post); err != nil {
		response := models.ErrorResponse{Error: "Failed to create post"}
		utils.ResponseWithJSON(w, response, http.StatusBadRequest)
		return
	}

	// Respond with success message
	response := models.SuccessResponse{Message: "Post Created Successfully"}
	utils.ResponseWithJSON(w, response, http.StatusOK)

}
