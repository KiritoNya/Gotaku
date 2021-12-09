package image

// Image is a struct that describes an image
type Image struct {
	//Id of image
	Id string

	//Url of image
	Url string

	//Name of image
	Name string

	//Data of image
	Data []byte

	// Format of image (.jpg,.png ecc...)
	Format string
}
