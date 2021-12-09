package anime

type Video struct {
	//Id: The id of the video
	Id string

	//Url: The url of the video
	Url string

	//Languages: The languages contained in the video
	Languages []string

	//Quality of the video (1080p, 720p, 480p...)
	Quality string

	//Duration: The duration of the video
	Duration int

	//Data: The data of the video
	Data []byte
}
