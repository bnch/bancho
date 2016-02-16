package packets

// ChannelJoin returns a packet of successful joining of a channel.
func ChannelJoin(channel string) Packet {
	return StringPacket(PacketChannelJoin, channel)
}
