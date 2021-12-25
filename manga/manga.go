package manga

import "github.com/KiritoNya/gotaku"

//Manga - Struct that describes the manga info
type Manga struct {
	gotaku.Media

	//NumChapters: The number of chapters
	NumChapters int

	//NumVolumes: The number of volumes
	NumVolumes int

	//Volumes: List of manga's volumes.
	Volumes []*Volume

	//Fansub: The fansub who translated the manga.
	Fansub Fansub

	//Keywords: The search keys to find the manga
	Keywords string
}
