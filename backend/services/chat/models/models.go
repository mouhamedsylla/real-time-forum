package models

import "real-time-forum/orm"

type Message struct {
	orm.Model
	SenderId   string `orm-go:"NOT NULL"`
	ReceiverId string `orm-go:"NOT NULL"`
	Content    string `orm-go:"NOT NULL"`
}

type Notification struct {
	UserID    string
	MessageId int
	Read      string
}

type Request struct {
	Token string `json:"token"`
}

type Response struct {
	Message string `json:"message"`
}

func NewMessage(sender, receiver, content string) Message {
	return Message{
		SenderId:   sender,
		ReceiverId: receiver,
		Content:    content,
	}
}

func NewNotification(messageId int, userId, read string) Notification {
	return Notification{
		UserID:    userId,
		MessageId: messageId,
		Read:      read,
	}
}

// notifications := []Notification{
// 	{Type: "new_message", Message: "You have a new message from {sender}."},
// 	{Type: "error", Message: "An error occurred. Please try again later."},
// 	{Type: "user_online", Message: "{username} is now online."},
// 	{Type: "user_offline", Message: "{username} is now offline."},
// 	{Type: "connection_lost", Message: "Connection lost. Reconnecting..."},
// 	{Type: "connection_restored", Message: "Connection restored."},
// 	{Type: "message_sent", Message: "Your message was sent successfully."},
// 	{Type: "message_failed", Message: "Failed to send your message. Please try again."},
// }
