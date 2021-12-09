package user

import (
	"KiritoNya/gotaku"
	"KiritoNya/gotaku/anime"
	"KiritoNya/gotaku/manga"
)

//Favourites - User's favourite anime, manga, characters, staff & studios
type Favourites struct {
	//Anime: Favourite anime
	Anime []*anime.Anime

	//Manga: Favourite manga
	Manga []*manga.Manga

	//Characters: Favourite characters
	Characters []*gotaku.Character

	//Staff: Favourite staff
	Staff []*gotaku.Staff

	//Studios: Favourite Studios
	Studios []*gotaku.Studio
}
