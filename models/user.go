package models

import (
	"github.com/bnch/bancho/packets"
)

// User Permissions
const (
	PermissionBanned = 1 << iota
	PermissionModerator
	PermissionAdmin
)

// User is an user on bancho.
type User struct {
	ID          int
	Username    string `sql:"size:20"`
	Permissions uint32
	Email       string
	Password    string
}

// IsBanned tells whether a user is banned.
func (u User) IsBanned() bool {
	if u.Permissions&PermissionBanned == 0 {
		return false
	}
	return true
}

// IsModerator tells whether a user is a moderator.
func (u User) IsModerator() bool {
	if u.Permissions&PermissionModerator == 0 {
		return false
	}
	return true
}

// IsAdmin tells whether a user is an admin.
func (u User) IsAdmin() bool {
	if u.Permissions&PermissionAdmin == 0 {
		return false
	}
	return true
}

// GetColour gets an user's colour in the chat.
func (u User) GetColour() byte {
	switch {
	case u.IsModerator():
		return packets.ColourMod
	case u.IsAdmin():
		return packets.ColourAdmin
	default:
		return packets.ColourSupporter
	}
}
