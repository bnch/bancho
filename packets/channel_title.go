package packets

import (
	"github.com/bnch/bancho/pid"
)

// ChannelTitle returns a packet with the name, description and number of users in a channel.
func ChannelTitle(channelName, channelDescription string, users uint16) Packet {
	return StringStringShort(pid.BanchoChannelAvailable, channelName, channelDescription, users)
}
