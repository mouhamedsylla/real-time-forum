package chat

// Endpoint to send a private message
type sendMessage struct{}

// Endpoint to retrieve private messages between two users
type getPrivateMessage struct{}

// Endpoint to retrieve the list of users with whom the user has exchanged private messagess
type getPrivateMessageUsers struct{}