package main

import (
	"real-time-forum/server"
	"real-time-forum/server/microservices"
	"real-time-forum/services/auth"
	"real-time-forum/services/chat"
	"real-time-forum/services/notification"
	"real-time-forum/services/posts"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	allServices := []microservices.Service{
		&auth.Auth{},
		&chat.Chat{},
		&posts.Publish{},
		&notification.Notification{},
	}
	server := server.NewServer(allServices...)
	server.StartServices()
}
