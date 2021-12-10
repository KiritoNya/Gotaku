package list

//Group - List group of anime or manga entries
type Group struct {
	//Entries: Media list entries
	Entries []*List

	//Name: Name of list group
	Name string

	//IsCustomList: Define if it's a custom list
	IsCustomList bool

	//IsSplitCompletedList: Define if it's a split of completed list
	IsSplitCompletedList bool

	//Status: Define the status of the list group
	Status Status
}
