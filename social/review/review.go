package review

import (
	"KiritoNya/gotaku"
	"time"
)

//Review - A Review that features in an anime or manga
type Review struct {
	//Id: The id of the review
	Id string

	//Url: The url for the review page on the website
	Url string

	//UserId: The id of the review's creator
	UserId string

	//MediaId: The id of the review's media
	MediaId string

	//MediaType: For which type of media the review is for
	MediaType gotaku.MediaType

	//Summary: A short summary of the review
	Summary string

	//Body: The main review body text
	Body string

	//Rating: The total user rating of the review
	Rating int

	//RatingAmount: The amount of user ratings of the review
	RatingAmount int

	//Score: The review score of the media
	Score int

	//Private: If the review is not yet publicly published and is only viewable by creator
	Private bool

	//CreatedAt: The time of the thread creation
	CreatedAt time.Time

	//UpdatedAt: The time of the thread last update
	UpdatedAt time.Time
}
