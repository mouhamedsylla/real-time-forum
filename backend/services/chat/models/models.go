package models

import "real-time-forum/orm"

type Message struct {
	orm.Model
	SenderId   int
	ReceiverId int
	Content    string `orm-go:"NOT NULL"`
}

type Notification struct {
	SenderId  int
	ReceiverId int
	Read      string
}

type UserContact struct {
	UsersId []int `json:"usersId"`
}

type Request struct {
	Token string `json:"token"`
}

type Response struct {
	Message string `json:"message"`
}

func NewMessage(sender, receiver int, content string) Message {
	return Message{
		SenderId:   sender,
		ReceiverId: receiver,
		Content:    content,
	}
}

func NewNotification(senderId, receiverId int, read string) Notification {
	return Notification{
		SenderId:  senderId,
		ReceiverId: receiverId,
		Read:      read,
	}
}
