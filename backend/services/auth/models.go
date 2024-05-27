package auth

import (
	"net/http"
	"real-time-forum/server/microservices"
)

type userLogin struct{
	orm.Model
	email string `orm-go:"NOT NULL"`
	password string `orm-go:"NOT NULL"`
}

type userRegister struct{
	orm.Model
	nickname string `orm-go:"NOT NULL"`
	age int `orm-go:"NOT NULL"`
	gender string `orm-go:"NOT NULL"`
	firstName string `orm-go:"NOT NULL"`
	lastName string `orm-go:"NOT NULL"`
	email string `orm-go:"NOT NULL"`
	password string `orm-go:"NOT NULL"`
}
