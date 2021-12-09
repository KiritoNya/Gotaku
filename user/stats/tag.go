package stats

import "KiritoNya/gotaku"

//Tag - Struct that describes the stat by tag
type Tag struct {
	stat
	Tag *gotaku.Tag
}
