package user

//NotificationType - Notification type enum
type NotificationType int

const (
	//ACTIVITY_MESSAGE - A user has sent you message
	ACTIVITY_MESSAGE NotificationType = iota

	//ACTIVITY_REPLY - A user has replied to your activity
	ACTIVITY_REPLY

	//FOLLOWING - A user has followed you
	FOLLOWING

	//ACTIVITY_MENTION - A user has mentioned you in their activity
	ACTIVITY_MENTION

	//THREAD_COMMENT_MENTION - A user has mentioned you in a forum comment
	THREAD_COMMENT_MENTION

	//THREAD_SUBSCRIBED - A user has commented in one of your subscribed forum threads
	THREAD_SUBSCRIBED

	//THREAD_COMMENT_REPLY - A user has replied to your forum comment
	THREAD_COMMENT_REPLY

	//AIRING - An anime you are currently watching has aired
	AIRING

	//ACTIVITY_LIKE - A user has liked your activity
	ACTIVITY_LIKE

	//ACTIVITY_REPLY_LIKE - A user has liked your activity reply
	ACTIVITY_REPLY_LIKE

	//THREAD_LIKE - A user has liked your forum thread
	THREAD_LIKE

	//THREAD_COMMENT_LIKE - A user has liked your forum comment
	THREAD_COMMENT_LIKE

	//ACTIVITY_REPLY_SUBSCRIBED - A user has replied to activity you have also replied to
	ACTIVITY_REPLY_SUBSCRIBED

	//RELATED_MEDIA_ADDITION - A new anime or manga has been added to the site where its related media is on the user's list
	RELATED_MEDIA_ADDITION

	//MEDIA_DATA_CHANGE - An anime or manga has had a data change that affects how a user may track it in their lists
	MEDIA_DATA_CHANGE

	//MEDIA_MERGE - Anime or manga entries on the user's list have been merged into a single entry
	MEDIA_MERGE

	//MEDIA_DELETION - An anime or manga on the user's list has been deleted from the site
	MEDIA_DELETION
)
