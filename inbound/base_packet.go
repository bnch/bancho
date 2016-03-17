package inbound

import (
	"bytes"

	"github.com/bnch/osubinary"
)

// BasePacket is the most basic type of packet that must then be converted into an actual packet.
type BasePacket struct {
	ID          uint16
	Content     []byte
	Initialised bool
}

// Unmarshal decodes some data from Content and puts it into an interface.
func (b BasePacket) Unmarshal(results ...interface{}) error {
	data := bytes.NewReader(b.Content)
	return osubinary.New(data).Unmarshal(results...)
}
