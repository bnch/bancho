package packets

import (
	"github.com/bnch/bancho/pid"
)

// ChannelJoin returns a packet of successful joining of a channel.
func ChannelJoin(channel string) Packet {
	return StringPacket(pid.BanchoChannelJoinSuccess, channel)
}
