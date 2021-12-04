package gotaku

type Status string

const (
	//Finished - Has completed and is no longer being released
	FINISHED Status = "Finished"

	//Releasing - Currently releasing
	RELEASING Status = "Releasing"

	//NotYetReleased - To be released at a later date
	NOT_YET_RELEASED Status = "Not Yet Released"

	//Cancelled - Ended before the work could be finished
	CANCELLED Status = "Cancelled"

	//Hiatus - Is currently paused from releasing and will resume at a later date
	HIATUS Status = "Hiatus"
)
