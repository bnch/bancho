package packets

import (
	"bytes"
	"encoding/binary"

	"github.com/bnch/bancho/pid"
)

// ChatMessage returns a packet yielding a chat message.
func ChatMessage(from, to, content string, userID int32) Packet {
	b := new(bytes.Buffer)

	b.Write(BanchoString(from))
	b.Write(BanchoString(content))
	b.Write(BanchoString(to))
	binary.Write(b, binary.LittleEndian, userID)

	endB := b.Bytes()
	return MakePacket(pid.BanchoSendMessage, uint32(len(endB)), endB)
}
