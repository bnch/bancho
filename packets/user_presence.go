package packets

import (
	"github.com/bnch/bancho/pid"
)

// UserPresence returns a packets which contains an user that has come online.
func UserPresence(userID int32) Packet {
	return MakePacket(pid.BanchoUserPresenceSingle, 4, userID)
}
