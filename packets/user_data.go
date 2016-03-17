package packets

import (
	"bytes"
	"encoding/binary"

	"github.com/bnch/bancho/pid"
)

// These are the various colours an user can have in the osu! chat.
const (
	ColourNormal    = 0
	ColourSupporter = 4
	ColourMod       = 6
	ColourAdmin     = 16
)

// UserDataInfo is a struct containing user data that can be passed to UserData()
// for making an user data packet. User data user data user data.
type UserDataInfo struct {
	ID         int32
	PlayerName string
	UTCOffset  byte
	Country    byte
	Colour     byte
	Longitude  float32
	Latitude   float32
	Rank       uint32
}

// UserData returns a packet containing the data of an user.
func UserData(u UserDataInfo) Packet {
	b := new(bytes.Buffer)

	binary.Write(b, binary.LittleEndian, u.ID)
	binary.Write(b, binary.LittleEndian, BanchoString(u.PlayerName))
	binary.Write(b, binary.LittleEndian, []byte{
		u.UTCOffset,
		u.Country,
		u.Colour,
	})
	binary.Write(b, binary.LittleEndian, []float32{
		u.Longitude,
		u.Latitude,
	})
	binary.Write(b, binary.LittleEndian, u.Rank)

	endB := b.Bytes()
	return MakePacket(pid.BanchoUserPresence, uint32(len(endB)), endB)
}
