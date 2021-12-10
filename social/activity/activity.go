package activity

//Activity - Struct that describe a general activity
type Activity struct {
	//Id: The id of the activity
	Id string

	//Url: The url for the activity page on the AniList website
	Url string

	//CreatorId: The user id of the activity's creator
	CreatorId string

	//Type: The type of activity
	Type Type

	//ReplyCount: The number of activity replies
	ReplyCount int

	//IsLocked: If the activity is locked and can receive replies
	IsLocked bool

	//IsSubscribed: If the currently authenticated user is subscribed to the activity
	IsSubscribed bool

	//LikeCount: The amount of likes the activity has
	LikeCount int

	//IsLiked: If the currently authenticated user liked the activity
	IsLiked int

	//CreatedAt: The time the activity was created at
	CreatedAt int

	//Replies: The written replies to the activity
	Replies []*Reply

	//Likes: The ids of users who liked the activity
	Likes []string
}
