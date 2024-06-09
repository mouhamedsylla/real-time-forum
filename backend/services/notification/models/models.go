package models

import "real-time-forum/orm"

type UserNotification struct {
	orm.Model
	UserID    string `orm-go:"NOT NULL"`
	MessageId int
	Read      string `orm-go:"NOT NULL"`
}

type NotifCount struct {
	Count int `json:"count"`
}

type Response struct {
	Message string `json:"message"`
}

type Request struct {
	Token string `json:"token"`
}
