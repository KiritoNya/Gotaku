package thread

import "time"

//Thread - Forum Thread
type Thread struct {
	//Id: The id of the thread
	Id string

	//Url: The url for the thread page on the AniList website
	Url string

	//Title: The title of the thread
	Title string

	//Body: The text body of the thread (Markdown)
	Body string

	//UserId: The id of the thread owner user
	UserId string

	//ReplyUserId: The id of the user who most recently commented on the thread
	ReplyUserId string

	//ReplyCommentId: The id of the most recent comment on the thread
	ReplyCommentId int

	//ReplyCount: The number of comments on the thread
	ReplyCount int

	//ViewCount: The number of times users have viewed the thread
	ViewCount int

	//IsLocked: If the thread is locked and can receive comments
	IsLocked bool

	//IsSticky: If the thread is stickied and should be displayed at the top of the page
	IsSticky bool

	//IsSubscribed: If the currently authenticated user is subscribed to the thread
	IsSubscribed bool

	//LikeCount: The amount of likes the thread has
	LikeCount int

	//IsLiked: If the currently authenticated user liked the thread
	IsLiked bool

	//RepliedAt: The time of the last reply
	RepliedAt time.Time

	//CreatedAt: The time of the thread creation
	CreatedAt time.Time

	//UpdatedAt: The time of the thread last update
	UpdatedAt time.Time

	//Likes: The users who liked the thread
	Likes []string

	//Categories: The categories of the thread
	Categories []*Category
}
