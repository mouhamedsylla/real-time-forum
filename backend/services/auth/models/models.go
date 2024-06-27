package models

import (
	"errors"
	"log"
	"real-time-forum/orm"

	"golang.org/x/crypto/bcrypt"
)

type Response struct {
	Message string `json:"message"`
}

type Request struct {
	Token string `json:"token"`
}

type LoggedUser struct {
	Id       int    `json:"id"`
	Nickname string `json:"nickname"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Email    string `json:"email"`
	Message  string `json:"message"`
}

type UserLogin struct {
	orm.Model
	Identifier string `json:"identifier"`
	Password   string `json:"password"`
}

type UserContact struct {
	UsersId []int `json:"usersId"`
}

type UserRegister struct {
	orm.Model
	Nickname  string `orm-go:"NOT NULL UNIQUE" json:"nickname" validate:"username"`
	Age       int    `orm-go:"NOT NULL" json:"age" validate:"min(18)"`
	Gender    string `orm-go:"NOT NULL" json:"gender" validate:"required"`
	FirstName string `orm-go:"NOT NULL" json:"firstName" validate:"required"`
	LastName  string `orm-go:"NOT NULL" json:"lastName" validate:"required"`
	Email     string `orm-go:"NOT NULL UNIQUE" json:"email" validate:"email"`
	Password  string `orm-go:"NOT NULL" json:"password" validate:"required"`
}

func Authenticate(Password string, toAuthenticate *UserLogin) error {
	err := bcrypt.CompareHashAndPassword([]byte(Password), []byte(toAuthenticate.Password))
	if err != nil {
		return errors.New("invalide email or password")
	}
	return nil
}

func CryptPassword(user *UserRegister) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
	if err != nil {
		log.Fatal(err)
	}
	user.Password = string(hashPassword)
}

func NewLoggedUser(user UserRegister, message string) LoggedUser {
	return LoggedUser{
		Id:       user.Id,
		Nickname: user.Nickname,
		FirstName: user.FirstName,
		LastName: user.LastName,
		Email:    user.Email,
		Message:  message,
	}
}
