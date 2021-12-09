package anime

import "time"

type AiringSchedule struct {
	//Id: The id of the airing schedule item
	Id int

	//AiringAt: The time the episode airs at
	AiringAt time.Time

	Episode *Episode
}
