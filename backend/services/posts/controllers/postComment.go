package controllers

import (
	"net/http"
	"real-time-forum/services/posts/database"
	"real-time-forum/services/posts/models"
	"real-time-forum/utils"
)

func (p *PostComment) HTTPServe() http.Handler {
	return http.HandlerFunc(p.PostComment)
}

func (p *PostComment) EndPoint() string {
	return "/posts/:postId/comment"
}

func (p *PostComment) SetMethods() []string {
	return []string{"POST"}
}

func (p *PostComment) PostComment(w http.ResponseWriter, r *http.Request) {
	CustomRoute := r.Context().Value("CustomRoute").(map[string]string)

	database.DbPost.Storage.Custom.Where("Id", CustomRoute["postId"])
	result := database.DbPost.Storage.Scan(models.UserPosts{}, "Id")
	
	if result == nil {
		utils.ResponseWithJSON(w, "Service Posts.postComment Post not found", http.StatusNotFound)
		return
	}
	com, status, err := utils.DecodeJSONRequestBody(r, models.Comments{})
	if err != nil {
		utils.ResponseWithJSON(w, err, status)
		return
	}

	comment := com.(*models.Comments)

	if err = database.DbPost.Storage.Insert(*comment); err != nil {
		utils.ResponseWithJSON(w, "Service Posts.postComment: 400 BadRequest", http.StatusBadRequest)
	}

	utils.ResponseWithJSON(w, "Comment posted Successfully", http.StatusOK)
}
