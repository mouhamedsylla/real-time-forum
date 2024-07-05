package models

import "real-time-forum/orm"

type UserPosts struct {
	orm.Model
	UserId     int
	Title      string `orm-go:"NOT NULL"`
	Image      []byte
	Content    string `orm-go:"NOT NULL"`
	Like       int    `json:"nbLike"`
	Dislike    int    `json:"nbDislike"`
	Categories string `orm-go:"NOT NULL"`
}

type Comments struct {
	orm.Model
	Comment string `orm-go:"NOT NULL"`
	Post_id int    `orm-go:"FOREIGN_KEY:Post:Id"`
	User_id int	
	Like    int
	Dislike int
}

type ReactionPost struct {
	Value  string `json:"value"`
	PostId int    `orm-go:"FOREIGN_KEY:Post:Id"`
	UserId int
}

type ReactionComment struct {
	Value     string `json:"value"`
	CommentId int    `orm-go:"FOREIGN_KEY:Comments:Id"`
	UserId    int
}

type LastCreated struct {
	LastId int64 `json:"lastId"`
}

type Response struct {
	Message string `json:"message"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
