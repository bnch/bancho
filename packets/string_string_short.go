package packets

import (
	"bytes"
	"encoding/binary"
)

// StringStringShort AKA Class20
func StringStringShort(packetID uint16, s1, s2 string, i uint16) Packet {
	b := new(bytes.Buffer)

	binary.Write(b, binary.LittleEndian, append(BanchoString(s1), BanchoString(s2)...))
	binary.Write(b, binary.LittleEndian, i)

	endB := b.Bytes()
	return MakePacket(packetID, uint32(len(endB)), endB)
}
