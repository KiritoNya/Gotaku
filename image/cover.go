package image

// Cover is a struct that describes the media covers
type Cover struct {
	//Id of the cover image
	Id string

	//ExtraLarge: The cover image url of the media at its largest size. If this size isn't available, large will be provided instead.
	ExtraLarge Image

	//Large: The cover image url of the media at a large size
	Large Image

	//Medium: The cover image url of the media at medium size
	Medium Image

	//Color: Average #hex color of cover image
	Color string
}
