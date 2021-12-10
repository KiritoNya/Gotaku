package anime

import (
	"KiritoNya/gotaku"
	"KiritoNya/gotaku/anime/aiiring"
)

type Anime struct {
	gotaku.Media

	NumEpisodes int
	Duration    int
	Trailers    []*Trailer

	//Keywords: The search keys to find the anime
	Keywords []string

	//NextAiringEpisode: The anime's next episode airing schedule
	NextAiringEpisode aiiring.Schedule
}
