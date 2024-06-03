package posts

import (
	"net/http"
	"real-time-forum/server/microservices"
	"real-time-forum/server/middleware"
)

type Publish struct {
	Post *microservices.Microservice
}

func (post *Publish) ConfigureEndpoint() {
	for _, controller := range post.Post.Controllers {
		post.Post.Router.Method(http.MethodGet).
						Middleware(middleware.LogRequest).
						Handler(
							controller.EndPoint(), 
							controller.HTTPServe(),
						)
	}
}

func (post *Publish) InitService() (err error){
	controllers := []microservices.Controller{
		// add controller ...
	}
	post.Post = microservices.NewMicroservice("Publish&Comments", ":8181")
	post.Post.Controllers = append(post.Post.Controllers, controllers...)
	return
}

func (post *Publish) GetService() *microservices.Microservice {
	return post.Post
}