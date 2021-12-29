package gotaku

//MediaType - The type of the Media
type MediaType string

const (
	//ANIME_TYPE - Japanese Anime
	ANIME_TYPE MediaType = "Anime"

	//MANGA_TYPE - Japanese comics
	MANGA_TYPE MediaType = "Manga"

	//MANHUA_TYPE Chinese comics
	MANHUA_TYPE MediaType = "Manhua"

	//MANHWA_TYPE Korean comics
	MANHWA_TYPE MediaType = "Manhwa"

	//THAI_TYPE Thailand comics
	THAI_TYPE MediaType = "Thai"

	//VIETNAMITA Vietnam comics
	VIETNAMITA MediaType = "Vietnamita"

	//DOUJINSHI_TYPE - Autopublicated japanese comics
	DOUJINSHI_TYPE MediaType = "Doujinshi"
)
