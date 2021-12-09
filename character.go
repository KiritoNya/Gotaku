package gotaku

import (
	"KiritoNya/gotaku/image"
	"time"
)

type Character struct {
	//Id: The id of the character
	Id int

	//Url: The url for the character page on the website
	Url string

	//Name: The names of the character
	Name string

	//Image: Character image
	Image image.Character

	//Description: A general description of the character
	Description string

	//Gender: The character's gender. Usually Male, Female, or Non-binary but can be any string.
	Gender string

	//Birthday: The character's birth date
	Birthday time.Time

	//Age: The character's age. Note this is a string, not an int, it may contain further text and additional ages.
	Age int

	//BloodType: The characters blood type
	BloodType string

	//Favourites: The amount of user's who have favourited the character
	Favourites int

	//Media:
	Media []Media
}
