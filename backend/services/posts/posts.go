package posts

import (
	"net/http"
	"real-time-forum/server/microservices"
)

type Posts struct {
	Post *microservices.Microservice
}

func (post *Posts) ConfigureEndpoint() {
	for _, controller := range post.Post.Controllers {
		post.Post.Router.Method(http.MethodGet).Handler(controller.EndPoint(), controller.HTTPServe())
	}
}

func (post *Posts) InitService() {
	controllers := []microservices.Controller{
		// add controller ...
	}
	post.Post = microservices.NewMicroservice("Realtime Chat", ":9090")
	post.Post.Controllers = append(post.Post.Controllers, controllers...)
}

func (post *Posts) GetService() *microservices.Microservice {
	return post.Post
}