package gotaku

import "time"

type Media struct {
	Id              string
	Title           *Title
	StartDate       time.Time
	EndDate         time.Time
	Season          Season
	Format          Format
	Status          Status
	IsAdult         bool
	Genres          []Genre
	Score           int
	Popularity      int
	Source          Source
	CountryOfOrigin int
}
