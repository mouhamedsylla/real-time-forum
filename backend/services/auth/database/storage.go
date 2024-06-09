package database

import (
	"real-time-forum/orm"
)

type AuthDB struct {
	Storage *orm.ORM
}

var Db = &AuthDB{}
