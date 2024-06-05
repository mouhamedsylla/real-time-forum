package posts

import (
	"real-time-forum/orm"
	"real-time-forum/server/microservices"
	"real-time-forum/server/middleware"
	"real-time-forum/utils"
)

const (
	DB_NAME = "Post.db"
	DB_PATH = "../../services/posts/database/"
)

var (
	storage *orm.ORM
)

type Publish struct {
	Post *microservices.Microservice
}

func (post *Publish) ConfigureEndpoint() {
	for _, controller := range post.Post.Controllers {
		post.Post.Router.Method(controller.SetMethods()...).
			Middleware(middleware.LogRequest).
			Handler(
				controller.EndPoint(),
				controller.HTTPServe(),
			)
	}
}

func (post *Publish) InitService() (err error) {
	storage, err = utils.InitStorage(DB_NAME, DB_PATH,
		Comments{},
		UserPosts{},
		Categories{},
		ReactionPost{},
		ReactionComment{},
	)
	controllers := []microservices.Controller{
		// add controller..
		&CreatedPost{},
		&GetPost{},
	}
	post.Post = microservices.NewMicroservice("Publish&Comments", ":8181")
	post.Post.Controllers = append(post.Post.Controllers, controllers...)
	return
}

func (post *Publish) GetService() *microservices.Microservice {
	return post.Post
}
