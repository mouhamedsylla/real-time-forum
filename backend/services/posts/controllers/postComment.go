package controllers

import (
	"net/http"
	"real-time-forum/services/posts/database"
	"real-time-forum/services/posts/models"
	"real-time-forum/utils"
	"strconv"
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

	postId, err := strconv.Atoi(CustomRoute["postId"])
	if err != nil {
		utils.ResponseWithJSON(w, "Service Posts.postComment: 400 BadRequest", http.StatusBadRequest)
		return
	}

	database.DbPost.Storage.Custom.Where("Id", postId)
	result := database.DbPost.Storage.Scan(models.UserPosts{}, "Id")
	database.DbPost.Storage.Custom.Clear()
	
	if result == nil {
		utils.ResponseWithJSON(w, "Service Posts.postComment Post not found", http.StatusNotFound)
		return
	}
	data, status, err := utils.DecodeJSONRequestBody(r, models.Comments{})
	if err != nil {
		utils.ResponseWithJSON(w, err, status)
		return
	}

	comment := data.(*models.Comments)
	comment.Post_id = postId
	if err = database.DbPost.Storage.Insert(*comment); err != nil {
		utils.ResponseWithJSON(w, "Service Posts.postComment: 400 BadRequest", http.StatusBadRequest)
	}

	utils.ResponseWithJSON(w, "Comment posted Successfully", http.StatusOK)
}
