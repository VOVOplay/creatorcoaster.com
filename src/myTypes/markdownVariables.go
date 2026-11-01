package myTypes

type ChannelVariable struct {
	Name string
	ID   uint64
}

// ---

type RoleUserType int

const (
	Role RoleUserType = iota
	User
)

type RoleUserVariable struct {
	Name string
	ID   uint64
	Type RoleUserType
}

// ---

// This is the whole part before the actual channel id
var ChannelPartialURL string = "https://discord.com/channels/1075932452842909806/"
