package packets

// Packet is a packet that must arrive at a user's client.
type Packet struct {
	Content []byte
}

// NewPacket generate a new packet.
func NewPacket(content []byte) Packet {
	return Packet{
		Content: content,
	}
}
