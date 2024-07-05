package controllers

import (
	"net/http"
	"real-time-forum/services/posts/database"
	"real-time-forum/services/posts/models"
	"real-time-forum/utils"
)

func (c *GetAllcomment) HTTPServe() http.Handler {
	return http.HandlerFunc(c.GetAllcomment)
}

func (c *GetAllcomment) EndPoint() string {
	return "/posts/getAllcomment"
}

func (c *GetAllcomment) SetMethods() []string {
	return []string{"GET"}
}

func (c *GetAllcomment) GetAllcomment(w http.ResponseWriter, r *http.Request) {
	var response models.Response
	data := database.DbPost.Storage.Scan(models.Comments{}, "Id", "CreatedAt", "Comment", "Post_id", "User_id", "Like", "Dislike")

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
