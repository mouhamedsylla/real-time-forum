package posts

import (
	"real-time-forum/server/microservices"
	"real-time-forum/server/middleware"
	"real-time-forum/services/posts/controllers"
	"real-time-forum/services/posts/database"
	"real-time-forum/services/posts/models"
	"real-time-forum/utils"
)

const (
	DB_NAME = "Post.db"
	DB_PATH = "../../services/posts/database/"
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
	database.DbPost.Storage, err = utils.InitStorage(DB_NAME, DB_PATH,
		models.Comments{},
		models.UserPosts{},
		models.Categories{},
		models.ReactionPost{},
		models.ReactionComment{},
	)
	controllers := []microservices.Controller{
		// add controller..
		&controllers.GetAllPost{},
		&controllers.PostComment{},
		&controllers.GetAllPost{},
		&controllers.CreatedPost{},
		&controllers.GetPost{},
	}

	post.Post = microservices.NewMicroservice("Publish&Comments", ":8181")
	post.Post.Controllers = append(post.Post.Controllers, controllers...)
	return
}

func (post *Publish) GetService() *microservices.Microservice {
	return post.Post
}
