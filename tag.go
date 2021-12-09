package gotaku

type Tag struct {
	//Id: The id of the tag
	Id int

	//Name: The name of the tag
	Name string

	//Description: A general description of the tag
	Description string

	//Category: The categories of tags this tag belongs to
	Category string

	//Rank: The relevance ranking of the tag out of the 100 for this media
	Rank int

	//IsSpoiler: If the tag is a spoiler for this media
	IsSpoiler bool

	//IsAdult: If the tag is only for adult 18+ media
	IsAdult bool
}
