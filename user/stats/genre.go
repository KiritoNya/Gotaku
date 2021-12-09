package stats

import "KiritoNya/gotaku"

//Genre - Struct that describes the stat by genre
type Genre struct {
	stat
	Genre *gotaku.Genre
}
