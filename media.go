package gotaku

import (
	"github.com/KiritoNya/gotaku/image"
	"time"
)

//Media - Anime or Manga
type Media struct {
	//Id: The id of the media
	Id string

	//Url: The url of the media website
	Url string

	//Title: The official titles of the media in various languages
	Title Title

	//Description: Short description of the media's story and characters
	Description string

	//StartDate: The first official release date of the media
	StartDate time.Time

	//EndDate: The last official release date of the media
	EndDate time.Time

	//Season: The season the media was initially released in
	Season Season

	//Format: The format the media was released in
	Format Format

	//Status: The current releasing status of the media
	Status Status

	//Type: The type of the media; anime or manga
	Type MediaType

	//IsAdult: If the media is intended only for 18+ adult audiences
	IsAdult bool

	//Genres: The genres of the media
	Genres []Genre

	//Score: The score of the media
	Score int

	//Popularity: The number of users with the media on their list
	Popularity int

	//Source: Source type the media was adapted from.
	Source Source

	//CountryOfOrigin: Where the media was created. (ISO 3166-1 alpha-2)
	CountryOfOrigin string

	//TwitterHashtag: Official Twitter hashtags for the media
	TwitterHashtag string

	//UpdatedAt: When the media's data was last updated
	UpdatedAt time.Time

	//CoverImage: The cover images of the media
	CoverImage image.Cover

	//BannerImage: The banner image of the media
	BannerImage image.Image

	//Synonyms: Alternative titles of the media
	Synonyms []string

	//AverageScore: A weighted average score of all the user's scores of the media
	AverageScore int

	//MeanScore: Mean score of all the user's scores of the media
	MeanScore int

	//Trending: The amount of related activity in the past hour
	Trending int

	//Favourites: The amount of user's who have favourited the media
	Favourites int

	//ExtternalLinks: External links to another site related to the media
	ExternalLinks []*ExternalLinks

	//Rankings: The ranking of the media in a particular time span and format compared to other media
	Rankings []*Rank

	//Tags: List of tags that describes elements and themes of the media
	Tags []*Tag

	//Relations: Other media in the same or connecting franchise
	Relations []*Media

	//Characters: The characters in the media
	Characters []*Character

	//Staff: The staff who produced the media
	Staff []*Staff

	//Studios: The companies who produced the media
	Studios []*Studio

	//Trends: The media's daily trend stats
	Trends []*Trend

	//Recommendations: User recommendations for similar media
	Recommendations Recommendation

	//Stats: The media's stats
	Stats Stats

	//Keywords: The search keys to find the media
	Keywords []string
}
