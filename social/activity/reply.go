package activity

import "time"

//Reply - Replay to an activity item
type Reply struct {
	//Id: The id of the reply
	Id string

	//CreatorId: The id of the replies creator
	CreatorId string

	//ActivityId: The id of the parent activity
	ActivityId string

	//Text: The reply text
	Text string

	//LikeCount: The amount of likes the reply has
	LikeCount int

	//IsLiked: If the currently authenticated user liked the reply
	IsLiked bool

	//CreatedAt: The time the reply was created at
	CreatedAt time.Time

	//Likes: The id of users who liked the reply
	Likes []string
}
