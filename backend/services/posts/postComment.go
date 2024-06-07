package posts

import (
	"net/http"
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

	storage.Custom.Where("Id", CustomRoute["postId"])

	result := storage.Scan(UserPosts{}, "Id")    
	if result == nil {
		utils.ResponseWithJSON(w, "Service Posts.postComment Post not found", http.StatusNotFound)
		return
	}
	// fmt.Println("PostCommentId: ", CustomRoute["postId"])

	com, status, err := utils.DecodeJSONRequestBody(r, Comments{})
	if err != nil {
		utils.ResponseWithJSON(w, err, status)
		return
	}

	comment := com.(*Comments)
		
	if err = storage.Insert(*comment); err != nil {
		utils.ResponseWithJSON(w, "Service Posts.postComment: 400 BadRequest", http.StatusBadRequest)
	}

	utils.ResponseWithJSON(w, "Comment posted Successfully", http.StatusOK)
}
