package packets

import (
	"bytes"
	"encoding/binary"

	"github.com/bnch/bancho/pid"
)

// UserQuit returns a packet about an user quitting.
func UserQuit(user int32) Packet {
	b := new(bytes.Buffer)
	binary.Write(b, binary.LittleEndian, user)
	binary.Write(b, binary.LittleEndian, byte(0))
	endB := b.Bytes()
	return MakePacket(pid.BanchoHandleUserQuit, uint32(len(endB)), endB)
}
