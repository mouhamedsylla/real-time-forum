package auth

import (
	"real-time-forum/orm"
)

type userLogin struct {
	orm.Model
	Identifier string `json:"identifier"`
	Password   string `json:"password"`
}

type userRegister struct {
	orm.Model
	Nickname  string `orm-go:"NOT NULL UNIQUE"`
	Age       int    `orm-go:"NOT NULL"`
	Gender    string `orm-go:"NOT NULL"`
	FirstName string `orm-go:"NOT NULL"`
	LastName  string `orm-go:"NOT NULL"`
	Email     string `orm-go:"NOT NULL UNIQUE"`
	Password  string `orm-go:"NOT NULL"`
}


