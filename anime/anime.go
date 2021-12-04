package anime

import "KiritoNya/gotaku"

type Anime struct {
	gotaku.Media

	NumEpisodes int
	Duration    int
}
