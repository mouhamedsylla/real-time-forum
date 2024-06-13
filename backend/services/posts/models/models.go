package models

import "real-time-forum/orm"

type UserPosts struct {
	orm.Model
	UserId string
	Title   string `orm-go:"NOT NULL"`
	Image 	[]byte
	Content string `orm-go:"NOT NULL"`
	Like    int    `json:"nbLike"`
	Dislike int    `json:"nbDislike"`
}

type Comments struct {
	orm.Model
	Comment string `orm-go:"NOT NULL"`
	Post_id int    `orm-go:"FOREIGN_KEY:Post:Id"`
	Like    int
	Dislike int
}

type Categories struct {
	Name    string `orm-go:"NOT NULL"`
	Id_Post int    `orm-go:"FOREIGN_KEY:Post:Id"`
}

type ReactionPost struct {
	Value  string `json:"value"`
	PostId int    `orm-go:"FOREIGN_KEY:Post:Id"`
}

type ReactionComment struct {
	Value     string `json:"value"`
	CommentId int    `orm-go:"FOREIGN_KEY:Comments:Id"`
}


type SuccessResponse struct {
	Message string `json:"message"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
