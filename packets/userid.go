package packets

import (
	"github.com/bnch/bancho/pid"
)

// These are the various error codes that can happen on login.
const (
	LoginFailed     = -1
	LoginNeedUpdate = -2
	// Locks the client
	LoginBanned = -4
	// LoginError is to be used in case of maintenance
	LoginError         = -5
	LoginNeedSupporter = -6
)

// UserID returns a packet containing the UserID. This is used on login to tell the client everything went smoothly.
//
// In case you wish to do other stuff on login, for instance locking the client, then use one of the constants beginning with `Login`.
func UserID(userID int32) Packet {
	return MakePacket(pid.BanchoLoginReply, 4, userID)
}
