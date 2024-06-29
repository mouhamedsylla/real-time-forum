package controllers

import (
	"net/http"
	"real-time-forum/services/posts/database"
	"real-time-forum/services/posts/models"
	"real-time-forum/utils"
)

func (p *GetPost) HTTPServe() http.Handler {
	return http.HandlerFunc(p.GetPost)
}

func (p *GetPost) EndPoint() string {
	return "/posts/:postId"
}

func (p *GetPost) SetMethods() []string {
	return []string{"GET"}
}

func (p *GetPost) GetPost(w http.ResponseWriter, r *http.Request) {
	CustomRoute := r.Context().Value("CustomRoute").(map[string]string)
	database.DbPost.Storage.Custom.Where("Id", CustomRoute["postId"])

	result := database.DbPost.Storage.Scan(models.UserPosts{}, "Id", "CreatedAt", "Title", "Content", "Like", "Dislike").([]models.UserPosts)
	if len(result) == 0 {
		utils.ResponseWithJSON(w, "Error Message from Posts.get_post: No Post Found", http.StatusNotFound)
		return
	}

	database.DbPost.Storage.Custom.Clear()
	utils.ResponseWithJSON(w, result[0], http.StatusOK)

}
