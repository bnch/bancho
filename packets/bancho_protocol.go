package packets

// ChoProtocol returns a packet with the current protocol version (which must be passed as an argument).
func ChoProtocol(protocolVersion uint32) Packet {
	return MakePacket(PacketChoProtocol, 4, protocolVersion)
}
