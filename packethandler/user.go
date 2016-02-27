package packethandler

import (
	"github.com/bnch/bancho/common"
	"github.com/bnch/bancho/packets"
)

// User represents an user online on bancho.
type User struct {
	ID     int32
	Name   string
	Token  string
	UTCOffset byte
	Country byte
	Colour byte
	Position struct {
		Longitude float32
		Latitude float32
	}
	Rank int
	Username string
	Status common.Status
}

// ToUserDataINfo converts a packethandler.User to a packets.UserDataInfo.
func (u User) ToUserDataInfo() packets.UserDataInfo {
	return packets.UserDataInfo{
			ID:         u.ID,
			PlayerName: u.Username,
			UTCOffset:  u.UTCOffset,
			Country:    u.Country,
			Colour:     u.Colour,
			Longitude:  u.Position.Longitude,
			Latitude:   u.Position.Latitude,
			Rank:       uint32(u.Rank),
		}
}
