package packets

// FriendList returns an int array of friend user IDs (passed as argument).
func FriendList(friends []int32) Packet {
	return IntArray(PacketFriendList, friends)
}
