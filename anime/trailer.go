package anime

import (
	"KiritoNya/gotaku/image"
)

//Trailer is a struct that describes the trailer infos
type Trailer struct {
	//Id: The trailer video id
	Id string

	//Site: The site the video is hosted by (Currently either youtube or dailymotion)
	Site string

	//Thumbnail: The url for the thumbnail image of the video
	Thumbnail *image.Image
}
