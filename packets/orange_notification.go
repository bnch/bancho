package packets

import (
	"github.com/bnch/bancho/pid"
)

// OrangeNotification returns a packet that will be displayed by the client as the "orange notification".
func OrangeNotification(message string) Packet {
	return StringPacket(pid.BanchoAnnounce, message)
}
