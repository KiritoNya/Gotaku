package list

//Status - Media list watching/reading status enum.
type Status int

const (
	//CURRENT - Currently watching/reading
	CURRENT Status = iota

	//PLANNING - Planning to watch/read
	PLANNING

	//COMPLETED - Finished watching/reading
	COMPLETED

	//DROPPED - Stopped watching/reading before completing
	DROPPED

	//PAUSED - Paused watching/reading
	PAUSED

	//REPEATING - Re-watching/reading
	REPEATING
)
