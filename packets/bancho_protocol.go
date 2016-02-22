package packets

import (
	"github.com/bnch/bancho/pid"
)

// ChoProtocol returns a packet with the current protocol version (which must be passed as an argument).
func ChoProtocol(protocolVersion uint32) Packet {
	return MakePacket(pid.BanchoProtocolVersion, 4, protocolVersion)
}
