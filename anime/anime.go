package anime

import (
	"github.com/KiritoNya/gotaku"
)

type Anime struct {
	gotaku.Media

	NumEpisodes int
	Duration    int
	Trailers    []*Trailer
	Episodes    []*Episode

	//Keywords: The search keys to find the anime
	Keywords []string
}
