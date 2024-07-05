package controllers

import (
	"io"
	"net/http"
	"real-time-forum/services/posts/database"
	"real-time-forum/services/posts/models"
	"real-time-forum/utils"
	"strconv"
)

func (p *CreatedPost) HTTPServe() http.Handler {
	return http.HandlerFunc(p.CreatedPost)
}

func (p *CreatedPost) EndPoint() string {
	return "/posts/createdpost/:userId"
}

func (p *CreatedPost) SetMethods() []string {
	return []string{"POST"}
}

func (p *CreatedPost) CreatedPost(w http.ResponseWriter, r *http.Request) {
	var response models.Response
	CustomRoute := r.Context().Value("CustomRoute").(map[string]string)
	var post = &models.UserPosts{}

	post.Categories = r.FormValue("categories")
	post.Title = r.FormValue("title")
	post.Content = r.FormValue("content")
	imageFile, _, err := r.FormFile("image")
	if err != nil {
		response.Message = "Unable to get image file"
		utils.ResponseWithJSON(w, response, http.StatusBadRequest)
		return
	}
	defer imageFile.Close()

	// Lire le fichier image
	post.Image, err = io.ReadAll(imageFile)
	if err != nil {
		response.Message = "Unable to read image file"
		utils.ResponseWithJSON(w, response, http.StatusInternalServerError)
		return
	}

	// Insert the post into the storage
	id, _ := strconv.Atoi(CustomRoute["userId"])
	post.UserId = id
	rsl, err := database.DbPost.Storage.Insert(*post)
	if err != nil {
		response.Message = "Failed to create post"
		utils.ResponseWithJSON(w, response, http.StatusBadRequest)
		return
	}

	// Respond with success message
	var postResponse models.LastCreated
	postResponse.LastId, _ = rsl[0].LastInsertId()
	utils.ResponseWithJSON(w, postResponse, http.StatusOK)

}
