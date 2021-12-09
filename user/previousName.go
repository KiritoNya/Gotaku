package user

import "time"

//PreviousName - A user's previous name
type PreviousName struct {
	//Name: A previous name of the user.
	Name string

	//CreatedAt: When the user first changed from this name.
	CreatedAt time.Time

	//UpdatedAt: When the user most recently changed from this name.
	UpdatedAt time.Time
}
