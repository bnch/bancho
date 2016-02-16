package packets

// NullPacket can be used to send packets with no data in them (well, a uint32(0) actually)
func NullPacket(packetID uint16) Packet {
	return MakePacket(packetID, 4, uint32(0))
}
