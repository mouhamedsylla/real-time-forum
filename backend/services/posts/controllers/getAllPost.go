package controllers

import (
	"fmt"
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
	var response models.Response
	result := database.DbPost.Storage.Scan(models.UserPosts{}, "Id", "CreatedAt", "UserId", "Title", "Image", "Content", "Like", "Dislike", "Categories").([]models.UserPosts)
	if len(result) == 0 {
		response.Message = "Any comment found for this post."
		fmt.Println("Any comment found for this post.")
		utils.ResponseWithJSON(w, response, http.StatusNoContent)
		return
	}

	utils.ResponseWithJSON(w, result, http.StatusOK)
}
