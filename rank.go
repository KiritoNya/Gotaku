package gotaku

type RankType string

const (
	RATED   RankType = "Rated"
	POPULAR RankType = "Popular"
)

//Rank - The ranking of a media in a particular time span and format compared to other media
type Rank struct {
	//Id: The id of the rank
	Id string

	//Rank: The numerical rank of the media
	Rank int

	//RankType: The type of ranking
	Type RankType

	//Format: The format the media is ranked within
	Format Format

	//Season: The season the media is ranked within
	Season Season

	//Year: The year the media is ranked within
	Year int

	//AllTime: If the ranking is based on all time instead of a season/year
	AllTime bool

	//Context: String that gives context to the ranking type and time span
	Context string
}
