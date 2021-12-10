package manga

import "KiritoNya/gotaku/image"

//Page - Struct that describes the page info
type Page struct {
	Id     string
	Number int
	Url    string
	Image  image.Image
}
