package thread

import "time"

type Comment struct {
	//Id: The id of the comment
	Id string

	//Url: The url for the comment page on the AniList website
	Url string

	//UserId: The user id of the comment's owner
	UserId string

	//ThreadId: The id of thread the comment belongs to
	ThreadId string

	//Comment: The text content of the comment (Markdown)
	Comment string

	//LikeCount: The amount of likes the comment has
	LikeCount int

	//IsLiked: If the currently authenticated user liked the comment
	IsLiked bool

	//CreatedAt: The time of the comments creation
	CreatedAt time.Time

	//UpdatedAt: The time of the comments last update
	UpdatedAt time.Time

	//Likes: The users who liked the comment
	Likes []string

	//ChildComments: The comment's child reply comments
	ChildComments []*Comment

	//IsLocked: If the comment tree is locked and may not receive replies or edits
	IsLocked bool
}
