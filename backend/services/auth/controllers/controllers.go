package controllers

import "real-time-forum/server/microservices"



type GetGroupUserDiscussion struct{}
type Register struct{}
type Login struct{}

var AuthClient microservices.Client