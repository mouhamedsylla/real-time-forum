package database

import (
	"real-time-forum/orm"
)

type ChatDB struct {
	Storage *orm.ORM
}

var DbChat = &ChatDB{}