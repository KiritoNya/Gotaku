package aiiring

import (
	"KiritoNya/gotaku/anime"
	"time"
)

//Schedule - Anime airing schedule.
type Schedule struct {
	//Id: The id of the airing schedule item
	Id int

	//AiringAt: The time the episode airs at
	AiringAt time.Time

	//Episode: The airing episode number
	Episode *anime.Episode

	//MediaId: The associate media id of the airing episode
	MediaId string
}
