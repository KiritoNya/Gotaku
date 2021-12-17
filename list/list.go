package list

import (
	"github.com/KiritoNya/gotaku"
	"time"
)

//List - List of anime or manga
type List struct {
	//Id: The id of the list entry
	Id string

	//UserId: The id of the user owner of the list entry
	UserId string

	//MediaId: The id of the media
	MediaId string

	//Status: The watching/reading status
	Status gotaku.Status

	//Score: The score of the entry
	Score float32

	//Progress: The amount of episodes/chapters consumed by the user
	Progress int

	//Repeat: The amount of times the user has rewatched/read the media
	Repeat int

	//Priority: Priority of planning
	Priority int

	//Private: If the entry should only be visible to authenticated user
	Private bool

	//HiddenFromStatusLists: If the entry shown be hidden from non-custom lists
	HiddenFromStatusLists bool

	//CustomLists: Map of booleans for which custom lists the entry are in
	CustomLists []CustomLists

	//AdvancedScores: Map of advanced scores with name keys
	AdvancedScores AdvancedScores

	//StartedAt: When the entry was started by the user
	StartedAt time.Time

	//CompletedAt: When the entry was completed by the user
	CompletedAt time.Time

	//UpdatedAt: When the entry data was last updated
	UpdatedAt time.Time

	//CreatedAt: When the entry data was created
	CreatedAt time.Time
}
