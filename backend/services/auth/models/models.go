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

type UserLogin struct {
	orm.Model
	Identifier string `json:"identifier"`
	Password   string `json:"password"`
}

type UserRegister struct {
	orm.Model
	Nickname  string `orm-go:"NOT NULL UNIQUE" json:"nickname"`
	Age       int    `orm-go:"NOT NULL" json:"age"`
	Gender    string `orm-go:"NOT NULL" json:"gender"`
	FirstName string `orm-go:"NOT NULL" json:"firstName"`
	LastName  string `orm-go:"NOT NULL" json:"lastName"`
	Email     string `orm-go:"NOT NULL UNIQUE" json:"email"`
	Password  string `orm-go:"NOT NULL" json:"password"`
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
