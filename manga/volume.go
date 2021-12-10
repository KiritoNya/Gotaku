package manga

//Volume - Struct that describe the volume info
type Volume struct {
	//Id: The volume id
	Id string

	//Url: The volume url
	Url string

	//Number: The volume number
	Number int

	//Chapters: The chapters contained in the volume
	Chapters []*Chapter
}
