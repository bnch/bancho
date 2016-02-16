package packets

// OnlinePlayers returns a packet with an array containing all the online users. User must be in it.
func OnlinePlayers(players []int32) Packet {
	return IntArray(PacketOnlinePlayers, players)
}
