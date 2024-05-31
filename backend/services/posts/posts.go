package posts

import (
	"net/http"
	"real-time-forum/server/microservices"
	"real-time-forum/server/middleware"
)

type Posts struct {
	Post *microservices.Microservice
}

func (post *Posts) ConfigureEndpoint() {
	for _, controller := range post.Post.Controllers {
		post.Post.Router.Method(http.MethodGet).Middleware(middleware.LogRequest).Handler(controller.EndPoint(), controller.HTTPServe())
	}
}

func (post *Posts) InitService() (err error){
	controllers := []microservices.Controller{
		// add controller ...
	}
	post.Post = microservices.NewMicroservice("Posts&Comments", ":8181")
	post.Post.Controllers = append(post.Post.Controllers, controllers...)
	return
}

func (post *Posts) GetService() *microservices.Microservice {
	return post.Post
}