package controllers

import (
	"fmt"
	"net/http"
	"real-time-forum/services/posts/database"
	"real-time-forum/services/posts/models"
	"real-time-forum/utils"
)

func (c *GetComment) HTTPServe() http.Handler {
	return http.HandlerFunc(c.GetComment)
}

func (c *GetComment) EndPoint() string {
	return "/posts/:postId/getcomment"
}

func (c *GetComment) SetMethods() []string {
	return []string{"GET"}
}

func (c *GetComment) GetComment(w http.ResponseWriter, r *http.Request) {
	CustomRoute := r.Context().Value("CustomRoute").(map[string]string)

	database.DbPost.Storage.Custom.Where("Post_id", CustomRoute["postId"])

	data := database.DbPost.Storage.Scan(models.Comments{}, "Id", "CreatedAt", "Comment", "Post_id", "Like", "Dislike")
	database.DbPost.Storage.Custom.Clear()

	if data == nil {
		fmt.Println("no comment")
		utils.ResponseWithJSON(w, "no comment", http.StatusNoContent)
		return
	}

	result := data.([]models.Comments)

	if len(result) == 0 {
		utils.ResponseWithJSON(w, "Error Message from Posts.postComment: No Comment Found", http.StatusNotFound)
		return
	}

	utils.ResponseWithJSON(w, result, http.StatusOK)
}
