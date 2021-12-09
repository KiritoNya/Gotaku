package anime

import "KiritoNya/gotaku"

type Anime struct {
	gotaku.Media

	NumEpisodes int
	Duration    int
	Trailer     Trailer

	//NextAiringEpisode: The anime's next episode airing schedule
	NextAiringEpisode AiringSchedule
}
