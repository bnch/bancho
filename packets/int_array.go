package packets

import (
	"bytes"
	"encoding/binary"
)

// IntArray generates a int[] without many clusterfucks.
func IntArray(packetID uint16, values []int32) Packet {
	b := new(bytes.Buffer)
	binary.Write(b, binary.LittleEndian, uint16(len(values)))
	binary.Write(b, binary.LittleEndian, values)
	endB := b.Bytes()
	return MakePacket(packetID, uint32(len(endB)), endB)
}
