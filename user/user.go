package user

import (
	"KiritoNya/gotaku/image"
	"KiritoNya/gotaku/user/stats"
	"time"
)

//User is a struct that describes the user info
type User struct {
	//Id: The id of the user
	Id int

	//Url: The url for the user page on the website
	Url string

	//Name: The name of the user
	Name string

	//About: The bio written by user
	About string

	//Avatar: The user's avatar images
	Avatar image.Avatar

	//Banner: The user's banner images
	Banner image.Image

	//Bans: The user ids of all bans users
	Bans []string

	//Options: The user's general options
	Options *Options

	//MediaListOptions: The user's media list options
	MediaListOptions MediaListOptions

	//Favourites: The users favourites
	Favourites Favourites

	//Statistics: The users anime & manga list statistics
	Statistics *stats.UserStatisticType

	//UnreadNotificationCount: The number of unread notifications the user has
	UnreadNotificationCount int

	//DonatorTier: The donation tier of the user
	DonatorTier int

	//DonationBadge: Custom donation badge text
	DonationBadge *image.Image

	//ModeratorRoles: The user's moderator roles if they are a site moderator
	ModeratorRoles []Role

	//CreatedAt: When the user's account was created. (Does not exist for accounts created before 2020)
	CreatedAt time.Time

	//UpdatedAt: When the user's data was last updated
	UpdatedAt time.Time

	PreviousName []*PreviousName
}
