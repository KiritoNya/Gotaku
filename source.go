package gotaku

type Source string

const (
	//ORIGINAL - An original production not based of another work
	ORIGINAL Source = "Original"

	//MANGA_SOURCE - Japanese comic book
	MANGA_SOURCE Source = "Manga"

	//LIGHT_NOVEL - Written work published in volumes
	LIGHT_NOVEL Source = "Light Novel"

	//VISUAL_NOVEL - Video game driven primary by text and narrative
	VISUAL_NOVEL Source = "Visual Novel"

	//VIDEO_GAME - Video game
	VIDEO_GAME Source = "Video Game"

	//DOUJINSHI - Self-published works
	DOUJINSHI Source = "Doujinshi"

	//ANIME - Japanese Anime
	ANIME Source = "Anime"

	//WEB_NOVEL - Written works published online
	WEB_NOVEL Source = "Web Novel"

	//LIVE_ACTION - Live action media such as movies or TV show
	LIVE_ACTION Source = "Live Action"

	//GAMES - Games excluding video games
	GAMES Source = "Games"

	//COMIC - Comics excluding manga
	COMIC Source = "Comic"

	//MULTIMEDIA_PROJECT - Multimedia project
	MULTIMEDIA_PROJECT Source = "Multimedia Project"

	//PICTURE_BOOK - Picture Book
	PICTURE_BOOK Source = "Picture Book"

	//OTHER - Other
	OTHER Source = "Other"
)
