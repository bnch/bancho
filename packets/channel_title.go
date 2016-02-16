package packets

// ChannelTitle returns a packet with the name, description and number of users in a channel.
func ChannelTitle(channelName, channelDescription string, users uint16) Packet {
	return StringStringShort(PacketChannelTitle, channelName, channelDescription, users)
}
