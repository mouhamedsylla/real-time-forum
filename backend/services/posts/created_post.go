package posts

import (
	"net/http"
	"real-time-forum/utils"
)

func (p *CreatedPost) HTTPServe() http.Handler {
	return http.HandlerFunc(p.CreatedPost)
}

func (p *CreatedPost) EndPoint() string {
	return "/posts/createdpost"
}

func (p *CreatedPost) SetMethods() []string {
	return []string{"POST"}
}

func (p *CreatedPost) CreatedPost(w http.ResponseWriter, r *http.Request) {
	pub, status, err := utils.DecodeJSONRequestBody(r, UserPosts{})
	if err != nil {
		utils.ResponseWithJSON(w, err, status)
		return
	}

	publi := pub.(*UserPosts)

	if err = storage.Insert(*publi); err != nil {
		utils.ResponseWithJSON(w, "Service posts.CreatedPost: 400 Bad Request", http.StatusBadRequest)
	}

	utils.ResponseWithJSON(w, "Post Created Successfully", http.StatusOK)

}
