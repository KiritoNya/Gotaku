package activity

//Type - Activity type enum.
type Type int

const (
	//TEXT - A text activity
	TEXT Type = iota

	//ANIME_LIST - A anime list update activity
	ANIME_LIST

	//MANGA_LIST - A manga list update activity
	MANGA_LIST

	//MESSAGE - A text message activity sent to another user
	MESSAGE

	//MEDIA_LIST - Anime & Manga list update, only used in query arguments
	MEDIA_LIST
)
