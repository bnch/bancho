package packethandler

import (
	"github.com/bnch/bancho/packets"
)

// Broadcast writes a packets to every user, except the one specified in "except".
func Broadcast(p packets.Packet, except string) {
	for token, sess := range Sessions {
		if token != except {
			sess.stream.Push(p)
		}
	}
}
