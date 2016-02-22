package packets

import (
	"github.com/bnch/bancho/pid"
)

// FriendList returns an int array of friend user IDs (passed as argument).
func FriendList(friends []int32) Packet {
	return IntArray(pid.BanchoFriendList, friends)
}
