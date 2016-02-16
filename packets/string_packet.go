package packets

// StringPacket returns a packet yielding a string.
func StringPacket(packetID uint16, s string) Packet {
	d := BanchoString(s)
	return MakePacket(packetID, uint32(len(d)), d)
}
