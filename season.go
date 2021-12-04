package gotaku

// SeasonPeriod is the season in which the media was released.
type SeasonPeriod string

const (
	WINTER SeasonPeriod = "Winter"
	SPRING SeasonPeriod = "Spring"
	SUMMER SeasonPeriod = "Summer"
	FALL   SeasonPeriod = "Fall"
)

// Season struct contains all info of season
type Season struct {
	Period SeasonPeriod
	Year   int
}
