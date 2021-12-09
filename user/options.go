package user

import "time"

type TitleLanguage string

const (
	//ROMAJI_TITLE - The romanization of the native language title
	ROMAJI_TITLE TitleLanguage = "Romaji"

	//ENGLISH_TITLE - The official english title
	ENGLISH_TITLE TitleLanguage = "English"

	//NATIVE_TITLE - Official title in it's native language
	NATIVE_TITLE TitleLanguage = "Native"
)

//StaffNameLanguage - The language the user wants to see staff and character names in
type StaffNameLanguage int

const (
	//ROMAJI_WESTERN - The romanization of the staff or character's native name, with western name ordering
	ROMAJI_WESTERN StaffNameLanguage = iota

	//ROMAJI - The romanization of the staff or character's native name
	ROMAJI

	//NATIVE - The staff or character's name in their native language
	NATIVE
)

//ScoreFormat - Media list scoring type
type ScoreFormat int

const (
	//POINT_100 - An integer from 0-100
	POINT_100 ScoreFormat = iota

	//POINT_10_DECIMAL - A float from 0-10 with 1 decimal place
	POINT_10_DECIMAL

	//POINT_10 - An integer from 0-10
	POINT_10

	//POINT_5 - An integer from 0-5. Should be represented in Stars
	POINT_5

	//POINT_3 - An integer from 0-3. Should be represented in Smileys. 0 => No Score, 1 => :(, 2 => :|, 3 => :)
	POINT_3
)

type Options struct {
	//TitleLanguage: The language the user wants to see media titles in
	TitleLanguage TitleLanguage

	//DisplayAdultContent: Whether the user has enabled viewing of 18+ content
	DisplayAdultContent bool

	//AiringNotifications: Whether the user receives notifications when a show they are watching aires
	AiiringNotifications bool

	//ProfileColor: Profile highlight color (hex)
	ProfileColor string

	//NotificationOptions: Notification options
	NotificationOptions NotificationOptions

	//Timezone: The user's timezone offset (Auth user only)
	Timezone time.Location

	//StaffNameLanguage: The language the user wants to see staff and character names in
	StaffNameLanguage StaffNameLanguage
}

//NotificationOptions - The notification options
type NotificationOptions struct {
	//Type: The type of notification
	Type NotificationType

	//Enabled: Whether this type of notification is enabled
	Enabled bool
}

//MediaListOptions - The struct describe a user's list options
type MediaListOptions struct {
	//ScoreFormat - The score format the user is using for media lists
	ScoreFormat ScoreFormat

	//RowOrder - The default order list rows should be displayed in
	RowOrder string

	AnimeList MediaListTypeOptions

	MangaList MediaListTypeOptions
}

//MediaListTypeOptions - A user's list options for anime or manga lists
type MediaListTypeOptions struct {
	//SectionOrder: The order each list should be displayed in
	SectionOrder []string

	//SplitCompletedSectionByFormat: If the completed sections of the list should be separated by format
	SplitCompletedSectionByFormat bool

	//CustomLists: The names of the user's custom lists
	CustomLists []string

	//AdvancedScoring: The names of the user's advanced scoring sections
	AdvancedScoring []string

	//AdvancedScoringEnabled: If advanced scoring is enabled
	AdvancedScoringEnabled bool
}
