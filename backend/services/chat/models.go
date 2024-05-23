package chat

import "real-time-forum/orm"

type Message struct {
	orm.Model
	SenderId   string `orm-go:"NOT NULL"`
	ReceiverId string `orm-go:"NOT NULL"`
	Content    string `orm-go:"NOT NULL"`
}

func NewMessage(sender, receiver, content string) Message {
	return Message{
		SenderId: sender,
		ReceiverId: receiver,
		Content: content,
	}
}

type Notification struct {
	Type    string `json:"type"`
	Message string `json:"message"`
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

