package packets

// P89 returns a Packet89. No-one has no idea of what packet 89 is, thus the generic name.
func P89() Packet {
	return NullPacket(Packet89)
}
