package gotaku

//Recommendation - Media recommendation
type Recommendation struct {
	//Id: The id of the recommendation
	Id string

	//Rating: Users rating of the recommendation
	Rating int

	//MediaId: The id of the media the recommendation came from
	MediaId int

	//MediaRaccomendationId: The id of the recommended media
	MediaRaccomendationId int

	//UserId: The user id that first created the recommendation
	UserId string
}
