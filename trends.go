package gotaku

import "time"

type Trend struct {
	//Id: The id of the trend
	Id string

	//Date: The day the data was recorded (timestamp)
	Date time.Time

	//Trending: The amount of media activity on the day
	Trending int

	//AverageScore: A weighted average score of all the user's scores of the media
	AverageScore int

	//Popularity: The number of users with the media on their list
	Popularity int

	//InProgress: The number of users with watching/reading the media
	InProgress int

	//Releasing: If the media was being released at this time
	Releasing bool

	//Episode: The episode number of the anime released on this day
	Episode int

	//Media: The related media
	Media Media
}
