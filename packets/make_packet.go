package packets

import (
	"bytes"
	"encoding/binary"
)

// MakePacket generates a new Packet with t being the type of packet and data being the actual data to write.
func MakePacket(t uint16, dataLen uint32, data interface{}) Packet {
	b := new(bytes.Buffer)
	binary.Write(b, binary.LittleEndian, t)
	binary.Write(b, binary.LittleEndian, byte(0))
	binary.Write(b, binary.LittleEndian, dataLen)
	binary.Write(b, binary.LittleEndian, data)
	p := Packet{
		Content: b.Bytes(),
	}
	return p
}
