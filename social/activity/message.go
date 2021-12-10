package activity

import "KiritoNya/gotaku/social/message"

//MessageActivity - User message activity
type MessageActivity struct {
	Activity
	message.Message
}
