package auth

import (
	"real-time-forum/orm"
)

type Response struct {
	Message string `json:"message"`
}

type Request struct {
	Token string `json:"token"`
}

type userLogin struct {
	orm.Model
	Identifier string `json:"identifier"`
	Password   string `json:"password"`
}

type UserRegister struct {
	orm.Model
	Nickname  string `orm-go:"NOT NULL UNIQUE"`
	Age       int    `orm-go:"NOT NULL"`
	Gender    string `orm-go:"NOT NULL"`
	FirstName string `orm-go:"NOT NULL"`
	LastName  string `orm-go:"NOT NULL"`
	Email     string `orm-go:"NOT NULL UNIQUE"`
	Password  string `orm-go:"NOT NULL"`
}


