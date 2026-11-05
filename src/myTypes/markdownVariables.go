package myTypes

// The type used in markdown articles for channel variables
type ChannelVariable struct {
	Name string // The name of the channel in discord
	ID   uint64 // The ID of the channel
}

// ---

type RoleUserType int

const (
	Role RoleUserType = iota
	User
)

// The type used in markdown articles for role-user variables
type RoleUserVariable struct {
	Name string       // The pretty name of the role/user (e.g. "VOVOplay", "Support")
	ID   uint64       // The ID of the role/user
	Type RoleUserType // Which type it is (myTypes.RoleUserType)
}

// ---

// This is the whole part before the actual channel id
var ChannelPartialURL string = "https://discord.com/channels/1075932452842909806/"
