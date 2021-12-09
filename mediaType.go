package gotaku

//MediaType - The type of the Media
type MediaType string

const (
	//ANIME_TYPE - Japanese Anime
	ANIME_TYPE MediaType = "Anime"

	//MANGA_TYPE - Japanese comics
	MANGA_TYPE MediaType = "Manga"

	//DOUJINSHI_TYPE - Autopublicated japanese comics
	DOUJINSHI_TYPE MediaType = "Doujinshi"
)
