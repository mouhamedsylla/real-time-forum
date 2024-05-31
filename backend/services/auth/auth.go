package auth

import (
	"errors"
	"log"
	"real-time-forum/orm"
	"real-time-forum/server/microservices"
	"real-time-forum/utils"

	"golang.org/x/crypto/bcrypt"
)

const (
	DB_NAME = "auth.db"
	DB_PATH = "../../services/auth/database/"
)

var storage *orm.ORM

type Auth struct {
	Auth *microservices.Microservice
}

func (auth *Auth) ConfigureEndpoint() {
	for _, controller := range auth.Auth.Controllers {
		auth.Auth.Router.Method(controller.SetMethods()...).Handler(
			controller.EndPoint(),
			controller.HTTPServe(),
		)
	}
}

func (auth *Auth) InitService() (err error){
	storage, err = utils.InitStorage(DB_NAME, DB_PATH, userRegister{}, userLogin{})
	controllers := []microservices.Controller{
		// add controller ...
		&Register{},
		&Login{},
	}
	auth.Auth = microservices.NewMicroservice("Authentication", ":8080")
	auth.Auth.Controllers = append(auth.Auth.Controllers, controllers...)
	return
}

func (auth *Auth) GetService() *microservices.Microservice {
	return auth.Auth
}

func Authenticate(Password string, toAuthenticate *userLogin) error {
	err := bcrypt.CompareHashAndPassword([]byte(Password), []byte(toAuthenticate.Password))
	if err != nil {
		return errors.New("invalide email or password")
	}
	return nil
}

func CryptPassword(user *userRegister) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
	if err != nil {
		log.Fatal(err)
	}
	user.Password = string(hashPassword)
}
