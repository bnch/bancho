package packethandler

import (
	"time"

	"github.com/bnch/bancho/inbound"
	"github.com/bnch/bancho/pid"
)

// RawPacketHandler handles inbound packets, and sends them to be analysed by functions.
func RawPacketHandler(pack inbound.BasePacket, s *Session) (delAfter bool) {
	if pack.ID == pid.OsuExit {
		UserQuit(s)
		delAfter = true
		return
	}
	go rawPacketHandler(pack, s)
	return
}
func rawPacketHandler(pack inbound.BasePacket, s *Session) {
	s.LastRequest = time.Now()
	switch pack.ID {
	case pid.OsuSendUserState:
		pack.Unmarshal(&s.User.Status.Status, &s.User.Status.Text, &s.User.Status.MapMD5)
	case pid.OsuSendIRCMessage:
		HandleMessage(pack, s)
	case pid.OsuChannelJoin:
		HandleChannelJoin(pack, s)
	case pid.OsuChannelLeave:
		HandleChannelPart(pack, s)
	}
}
