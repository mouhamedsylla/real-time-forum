package posts

import (
	"net/http"
)

func (p *GetPost) HTTPServe() http.Handler {
	return http.HandlerFunc(p.GetPost)
}

func (p *GetPost) Endpoint() string {
	return "/posts/:postId"
}

func (p *GetPost) SetMethods() []string {
	return []string{"GET"}
}

func (p *GetPost) GetPost(w http.ResponseWriter, r *http.Request) {

}
