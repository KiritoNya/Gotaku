package anime

import "KiritoNya/gotaku/language"

//Episode - The struct describe the episode info
type Episode struct {
	//Id: The id of the episode
	Id string

	//Url: The website url of the episode
	Url string

	//Num: The number of the episode
	Num string

	//Name: The title of the episode with the primary language
	Name string

	//AlternativeTitles: Other titles with different language
	AlternativeTitles map[language.Language]string

	//Thumbnail: The url of episode image thumbnail
	Thumbnail string

	//Site: The site location of the streaming episodes
	Site string

	//Videos: The episode videos
	Videos []*Video
}
