package gotaku

import (
	"github.com/KiritoNya/gotaku/image"
	"time"
)

//Staff is a struct that describes the staff info
type Staff struct {
	//Id: The id of the staff member
	Id string

	//Url: The url for the staff page on the website
	Url string

	//Name: The names of the staff member
	Name StaffName

	//Language: The primary language of the staff member. Current values: Japanese, English, Korean, Italian, Spanish, Portuguese, French, German, Hebrew, Hungarian, Chinese, Arabic, Filipino, Catalan
	Language string

	//Image: The staff images
	Image StaffImage

	//Description: A general description of the staff member
	Description string

	//PrimaryOccupations: The person's primary occupations
	PrimaryOccupations []string

	//Gender: The staff's gender. Usually Male, Female, or Non-binary but can be any string.
	Gender string

	//DateOfBirth: The date of birthday
	DateOfBirth time.Time

	//DateOfDeath: The date of death
	DateOfDeath time.Time

	//Age: The person's age in years
	Age int

	//YearsActive: [startYear, endYear] (If the 2nd value is not present staff is still active)
	YearsActive []int

	//HomeTown: The persons birthplace or hometown
	HomeTown string

	//BloodType: The persons blood type
	BloodType string

	//Role: The staff role
	Role string

	//StaffMedia: Media where the staff member has a production role
	StaffMedia []Media

	//Characters: Characters voiced by the actor
	Characters []Character

	//CharacterMedia: Media the actor voiced characters in. (Same data as characters with media as node instead of characters)
	CharacterMedia []Media
}

//StaffName is a struct that describes the name of Staff
type StaffName struct {
	//First: The person's given name
	First string

	//Middle: The person's middle name
	Middle string

	//Last: The person's surname
	Last string

	//Full: The person's first and last name
	Full string

	//Native: The person's full name in their native language
	Native string

	//Alternative: Other names the staff member might be referred to as (pen names)
	Alternative []string
}

// StaffImage is struct the described a staff image
type StaffImage struct {
	Id     string
	Large  *image.Image
	Medium *image.Image
}
