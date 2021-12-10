package anime

import "time"

//AiringSchedule - Media Airing Schedule. NOTE: We only aim to guarantee that FUTURE airing data is present and accurate.
type AiringSchedule struct {
	//Id: The id of the airing schedule item
	Id int

	//AiringAt: The time the episode airs at
	AiringAt time.Time

	//Episode: The airing episode number
	Episode *Episode

	//MediaId: The associate media id of the airing episode
	MediaId string
}
