package activity

import "github.com/KiritoNya/gotaku/social/message"

//MessageActivity - User message activity
type MessageActivity struct {
	Activity
	message.Message
}
