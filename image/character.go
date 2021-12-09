package image

// Character is a struct that describes the characters of media
type Character struct {
	//Id of the cover image
	Id string

	//Large: The character's image of media at its largest size
	Large Image

	//Medium: The character's image of media at medium size
	Medium Image
}
