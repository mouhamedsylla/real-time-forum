package posts

import (
	"net/http"
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

	result := storage.Scan(UserPosts{}, "Id", "CreatedAt", "Title", "Content", "Like", "Dislike").([]UserPosts)
	if len(result) == 0 {
		utils.ResponseWithJSON(w, "Error Message from Posts.getAllPost: No Post Found", http.StatusNotFound)
		return
	}

	utils.ResponseWithJSON(w, result, http.StatusOK)
}
