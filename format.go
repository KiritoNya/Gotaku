package gotaku

// Format is media format type
type Format string

const (
	//TV - Anime broadcast on television
	TV Format = "Tv"

	//TV_SHORT - Anime which are under 15 minutes in length and broadcast on television
	TV_SHORT Format = "Tv Short"

	//MOVIE - Anime movies with a theatrical release
	MOVIE Format = "Movie"

	//SPECIAL - episodes that have been included in DVD/Blu-ray releases, picture dramas, pilots, etc
	SPECIAL Format = "Special"

	//OVA - (Original Video Animation) Anime that have been released directly on DVD/Blu-ray without originally going through a theatrical release or television broadcast
	OVA Format = "OVA"

	//ONA - (Original Net Animation) Anime that have been originally released online or are only available through streaming services.
	ONA Format = "ONA"

	//MUSIC - Short anime released as a music video
	MUSIC Format = "Music"

	//MANGA - Professionally published manga with more than one chapter
	MANGA Format = "Manga"

	//NOVEL - Written books released as a series of light novels
	NOVEL Format = "Novel"

	//ONE_SHOT - Manga with just one chapter
	ONE_SHOT Format = "One Shot"
)
