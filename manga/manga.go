package manga

import "KiritoNya/gotaku"

type Manga struct {
	gotaku.Media

	NumChapters int
	NumVolumes  int
}
