package packethandler

import (
	"time"

	"github.com/bnch/bancho/inbound"
	"github.com/bnch/bancho/pid"
)

type packSess struct {
	p inbound.BasePacket
	s *Session
}

// RawPacketHandler handles inbound packets, and sends them to be analysed by functions.
func RawPacketHandler(pack inbound.BasePacket, s *Session) (delAfter bool) {
	ps := packSess{
		p: pack,
		s: s,
	}
	if pack.ID == pid.OsuExit {
		UserQuit(ps)
		delAfter = true
		return
	}
	go rawPacketHandler(ps)
	return
}
func rawPacketHandler(ps packSess) {
	ps.s.LastRequest = time.Now()
	switch ps.p.ID {
	case pid.OsuSendUserState:
		ps.p.Unmarshal(&ps.s.User.Status.Status, &ps.s.User.Status.Text, &ps.s.User.Status.MapMD5)
	case pid.OsuSendIRCMessage:
		HandleMessage(ps)
	case pid.OsuChannelJoin:
		HandleChannelJoin(ps)
	case pid.OsuChannelLeave:
		HandleChannelPart(ps)
	case pid.OsuUserStatsRequest:
		HandleUserStatsRequest(ps)
	}
}
