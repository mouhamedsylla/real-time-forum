package posts

import (
	"net/http"
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
	//method Aziz
	// urlpath := r.URL.Path
	// tab_url := strings.Split(urlpath, "/")
	// urlpath = tab_url[len(tab_url)-1]
	// fmt.Println("GetPostId: ", urlpath)

	// method mouhamed
	CustomRoute := r.Context().Value("CustomRoute").(map[string]string)
	// fmt.Println(CustomRoute["postId"])

	storage.Custom.Where("Id", CustomRoute["postId"])

	result := storage.Scan(UserPosts{}, "Id", "CreatedAt", "Title", "Content", "Like", "Dislike").([]UserPosts)
	if len(result) == 0 {
		utils.ResponseWithJSON(w, "Error Message from Posts.get_post: No Post Found", http.StatusNotFound)
		return
	}

	storage.Custom.Clear()

	utils.ResponseWithJSON(w, result[0], http.StatusOK)

}
