package gotaku

//Stats - Struct that describe a media's statistics
type Stats struct {
	ScoreDistribution  []ScoreDistribution
	StatusDistribution []StatusDistribution
}

//ScoreDistribution - Struct that describes a user's list score distribution.
type ScoreDistribution struct {
	//Score: The score
	Score int

	//Amount: The amount of list entries with this score
	Amount int
}

//StatusDistribution - The distribution of the watching/reading status of media or a user's list.
type StatusDistribution struct {
	//Status: The status
	Status Status

	//Amount: The amount of entries with this status
	Amount int
}
