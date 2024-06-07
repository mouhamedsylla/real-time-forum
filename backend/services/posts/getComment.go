package posts

import (
	"fmt"
	"net/http"
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

	storage.Custom.Where("Post_id", CustomRoute["postId"])
	
	result := storage.Scan(Comments{}, "Id", "CreatedAt", "Comment", "Post_id", "Like", "Dislike").([]Comments)

	if len(result) == 0 {
		utils.ResponseWithJSON(w, "Error Message from Posts.postComment: No Comment Found", http.StatusNotFound)
		return
	}

	fmt.Println("result: ", result)
	fmt.Println("GetCommentId: ", CustomRoute["postId"])

	storage.Custom.Clear()

	utils.ResponseWithJSON(w, result, http.StatusOK)
}
