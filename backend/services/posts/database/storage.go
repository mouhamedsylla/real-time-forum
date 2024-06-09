package database

import (
	"real-time-forum/orm"
)

type PostDB struct {
	Storage *orm.ORM
}

var DbPost = &PostDB{}