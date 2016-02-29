package packets

import (
	"bytes"
	"encoding/binary"
	"github.com/bnch/bancho/pid"
)

// UserDataFullInfo is a struct containing all the user data you'll ever need.
type UserDataFullInfo struct {
	ID int32
	// Action is afk, modding, playing, listening...
	Action       byte
	ActionText   string
	ActionMapMD5 string
	Mods         int32
	GameMode     byte
	// Nobody knows the fuck this is.
	// ripple bancho v1 says this should be "\x00\x00\x00\x01", although justm3's
	// custom-bancho says it should be just 0 (so "\x00\x00\x00\x00").
	// Beatmap /b/ ID?
	UnknownInt int32
	Score      uint64
	// Accuracy is divided by 100 when sent to osu!. This is because 0.1337 = 13.37% on the osu! client.
	Accuracy  float32
	Playcount uint32
	// Used for level
	TotalScore uint64
	Rank       uint32
	PP         uint16
}

// UserDataFull returns a packet containing the [extended] data of an user.
func UserDataFull(u UserDataFullInfo) Packet {
	b := new(bytes.Buffer)

	binary.Write(b, binary.LittleEndian, u.ID)
	binary.Write(b, binary.LittleEndian, u.Action)
	binary.Write(b, binary.LittleEndian, append(BanchoString(u.ActionText), BanchoString(u.ActionMapMD5)...))
	binary.Write(b, binary.LittleEndian, u.Mods)
	binary.Write(b, binary.LittleEndian, u.GameMode)
	binary.Write(b, binary.LittleEndian, u.UnknownInt)
	binary.Write(b, binary.LittleEndian, u.Score)
	binary.Write(b, binary.LittleEndian, u.Accuracy/100)
	binary.Write(b, binary.LittleEndian, u.Playcount)
	binary.Write(b, binary.LittleEndian, u.TotalScore)
	binary.Write(b, binary.LittleEndian, u.Rank)
	binary.Write(b, binary.LittleEndian, u.PP)

	endB := b.Bytes()
	return MakePacket(pid.BanchoHandleUserUpdate, uint32(len(endB)), endB)
}
