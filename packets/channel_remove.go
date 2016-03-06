package packets

import (
	"github.com/bnch/bancho/pid"
)

// ChannelRemove returns a packet that will make the client remove the channel.
func ChannelRemove(ch string) Packet {
	return StringPacket(pid.BanchoChannelRevoked, ch)
}
