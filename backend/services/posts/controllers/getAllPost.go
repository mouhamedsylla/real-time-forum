package controllers

import (
	"net/http"
	"real-time-forum/services/posts/database"
	"real-time-forum/services/posts/models"
	"real-time-forum/utils"
)

func (p *GetAllPost) HTTPServe() http.Handler {
	return http.HandlerFunc(p.GetAllPost)
}

func (p *GetAllPost) EndPoint() string {
	return "/posts/getAllPost"
}

func (p *GetAllPost) SetMethods() []string {
	return []string{"GET"}
}

func (p *GetAllPost) GetAllPost(w http.ResponseWriter, r *http.Request) {

	result := database.DbPost.Storage.Scan(models.UserPosts{}, "Id", "CreatedAt", "Title", "Image", "Content", "Like", "Dislike").([]models.UserPosts)
	if len(result) == 0 {
		utils.ResponseWithJSON(w, "Error Message from Posts.getAllPost: No Post Found", http.StatusNoContent)
		return
	}

	utils.ResponseWithJSON(w, result, http.StatusOK)
}
