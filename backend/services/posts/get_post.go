package posts

import (
	"fmt"
	"net/http"
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
	fmt.Println(CustomRoute["postId"])

	result := storage.Scan(UserPosts{}, "Id", "CreatedAt", "Title", "Content", "Like", "Dislike").([]UserPosts)
	fmt.Println(result)
}
