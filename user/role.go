package user

type Role int

const (
	//ADMIN - An administrator
	ADMIN Role = iota

	//LEAD_DEVELOPER - A head developer
	LEAD_DEVELOPER

	//DEVELOPER - An developer
	DEVELOPER

	//LEAD_COMMUNITY - A lead community moderator
	LEAD_COMMUNITY

	//COMMUNITY - A community moderator
	COMMUNITY

	//DISCORD_COMMUNITY - A discord community moderator
	DISCORD_COMMUNITY

	//LEAD_ANIME_DATA - An anime data moderator
	LEAD_ANIME_DATA

	//ANIME_DATA - An anime data moderator
	ANIME_DATA

	//LEAD_MANGA_DATA - A lead manga data moderator
	LEAD_MANGA_DATA

	//MANGA_DATA - A manga data moderator
	MANGA_DATA

	//LEAD_SOCIAL_MEDIA - A lead social media moderator
	LEAD_SOCIAL_MEDIA

	//SOCIAL_MEDIA - A social media moderator
	SOCIAL_MEDIA

	//RETIRED - A retired moderator
	RETIRED
)
