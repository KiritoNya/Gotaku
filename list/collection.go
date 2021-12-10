package list

//Collection - List of anime or manga
type Collection struct {
	//Lists: Grouped media list entries
	Lists []*Group

	//UserId: The owner of the list
	UserId int

	//HasNextChunk: If there is another chunk
	HasNextChunk bool
}
