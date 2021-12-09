package gotaku

//Studio is a struct that describes the studio info
type Studio struct {
	//Id: The id of the studio
	Id string

	//Url: The url for the studio page on the website
	Url string

	//Name: The name of the studio
	Name string

	//IsAnimationStudio: If the studio is an animation studio or a different kind of company
	IsAnimationStudio bool

	//Media: The media the studio has worked on
	Media []Media
}
