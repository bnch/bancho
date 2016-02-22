package packets

import (
	"github.com/bnch/bancho/pid"
)

// ChannelListingComplete tells bancho we finished outputting all the channels.
func ChannelListingComplete() Packet {
	return NullPacket(pid.BanchoChannelListingComplete)
}
