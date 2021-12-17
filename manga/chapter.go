package manga

import (
	"github.com/KiritoNya/gotaku"
	"time"
)

//Chapter - Struct that describes a manga chapter
type Chapter struct {
	//Id: The chapter id
	Id string

	//Url: The chapter url
	Url string

	//Number: The chapter number
	Number int

	//Title: The title of chapters
	Title gotaku.Title

	//Volume: The number of the volume in which the chapter is located.
	Volume int

	DateAdd time.Time

	//Pages: The chapters pages
	Pages []*Page

	//Keywords: The search keys to find the chapter
	Keywords []string
}
