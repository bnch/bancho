package packets

// SilenceClient makes a packet telling the client how much time the user has to be silence.
// time is the number of seconds until the end of the silence.
func SilenceClient(time uint32) Packet {
	return MakePacket(PacketSilenceClient, 4, time)
}
