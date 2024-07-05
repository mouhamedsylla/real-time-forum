package controllers

import (
	"net/http"
	"real-time-forum/services/posts/database"
	"real-time-forum/services/posts/models"
	"real-time-forum/utils"
	"strconv"
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
	var response models.Response
	CustomRoute := r.Context().Value("CustomRoute").(map[string]string)
	postId, _ := strconv.Atoi(CustomRoute["postId"])

	database.DbPost.Storage.Custom.Where("Post_id", postId)

	data := database.DbPost.Storage.Scan(models.Comments{}, "Id", "CreatedAt", "Comment", "Post_id", "Like", "Dislike")
	database.DbPost.Storage.Custom.Clear()

	if data == nil {
		response.Message = "no comment"
		utils.ResponseWithJSON(w, response, http.StatusNotFound)
		return
	}

	result := data.([]models.Comments)

	if len(result) == 0 {
		response.Message = "No Comment Found"
		utils.ResponseWithJSON(w, response, http.StatusNotFound)
		return
	}

	utils.ResponseWithJSON(w, result, http.StatusOK)
}
