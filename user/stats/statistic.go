package stats

//stat - struct that described base statistic infos
type stat struct {
	Count          int
	MeanScore      int
	MediaIds       []int
	ChaptersRead   int
	MinutesWatched int
}

//Statistics - Struct that describe the User stats
type Statistics struct {
	Count             int
	MeanScore         float32
	StandardDeviation float32
	Formats           []*Format
	Statuses          []*Status
	Scores            []*Score
	Lengths           []*Length
	ReleaseYears      []*ReleaseYear
	StartYears        []*StartYear
	Genres            []*Genre
	Tags              []*Tag
	Countries         []*Country
	VoiceActors       []*VoiceActor
	Staff             []*Staff
	Studios           []*Studio
}

//AnimeStatistics - Struct that described the Anime user stats
type AnimeStatistics struct {
	Statistics
	MinutesWatched  int
	EpisodesWatched int
}

//MangaStatistics - Struct that described the Manga user stats
type MangaStatistics struct {
	Statistics
	ChaptersRead int
	VolumesRead  int
}

//UserStatisticType - Struct that describes the user statistics
type UserStatisticType struct {
	Anime AnimeStatistics
	Manga MangaStatistics
}
