package database

import (
	"real-time-forum/orm"
)

type NotificationDB struct {
	Storage *orm.ORM
}

var DbNotification = &NotificationDB{}