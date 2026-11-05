package myTypes

// The struct used for the About Us page
type StaffMember struct {
	UserID             uint64 // ID of the user
	Username           string // Discord username (not display name)
	ProfilePictureLink string // The link to a user's pfp
	Position           string // The programmatic name of their positon (e.g. "head_admin")
	PositionPrettyName string // The pretty name of their position (e.g. "Head Admin")
}
