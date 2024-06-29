package models

import "real-time-forum/orm"

type UserNotification struct {
	orm.Model
	SenderId  int `json:"senderId"`
	ReceiverId int `json:"receiverId"`
	Read      string `orm-go:"NOT NULL"`
}

type ConnectedUser struct {
	Id       int    `json:"id"`
}

type UserInfos struct {
	Type	 string `json:"type"`
	Id       int    `json:"id"`
	Status   string `json:"status"`
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
