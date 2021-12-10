package message

type Message struct {
	//Id: The id of the activity
	Id string

	//SenderId: The user id of the sender
	SenderId string

	//RecipientsId: The user id of the recipients
	RecipientsId []string

	//Message: The message text
	Message string

	//IsPrivate: If the message is private and only viewable to the sender and recipients
	IsPrivate bool
}
