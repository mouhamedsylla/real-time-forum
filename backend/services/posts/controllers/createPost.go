package controllers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"real-time-forum/services/posts/database"
	"real-time-forum/services/posts/models"
	"real-time-forum/utils"
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
	CustomRoute := r.Context().Value("CustomRoute").(map[string]string)
	var post = &models.UserPosts{}

	post.Title = r.FormValue("title")
	post.Content = r.FormValue("content")
	imageFile, _, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "Unable to get image file", http.StatusBadRequest)
		return
	}
	defer imageFile.Close()

	// Lire le fichier image
	post.Image, err = ioutil.ReadAll(imageFile)
	if err != nil {
		http.Error(w, "Unable to read image file", http.StatusInternalServerError)
		return
	}

	// Insert the post into the storage
	post.UserId = CustomRoute["userId"]
	fmt.Println(CustomRoute)
	if err = database.DbPost.Storage.Insert(*post); err != nil {
		response := models.ErrorResponse{Error: "Failed to create post"}
		utils.ResponseWithJSON(w, response, http.StatusBadRequest)
		return
	}

	// Respond with success message
	response := models.SuccessResponse{Message: "Post Created Successfully"}
	utils.ResponseWithJSON(w, response, http.StatusOK)

}
