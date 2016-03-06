package packethandler

import (
	"time"

	"github.com/bnch/bancho/inbound"
	"github.com/bnch/bancho/pid"
)

// RawPacketHandler handles inbound packets, and sends them to be analysed by functions.
func RawPacketHandler(pack inbound.BasePacket, s *Session) (delAfter bool) {
	s.LastRequest = time.Now()
	switch pack.ID {
	case pid.OsuSendUserState:
		/* E.G.:
		* 00 00 00 55 00 00 00 06 0b 27 6f 73 75 21 20 70 6c 61 79 20 4f 6d 6f 69 20 2d 20 53 6e 6f 77 20 44 72 69 76 65 20 5b 53 6e 6f 77 44 72 69 76 65 5d
		* ^____________ header ^  ^_____________________________________________________________________________________________________________ Action text
		*                      |_ Status (common.Status constants)
		* 0b 20 36 34 39 32 30 34 33 61 33 31 33 61 37 34 30 61 61 36 38 32 66 63 61 63 64 38 64 66 36 38 38 31 00 08 00 00 00
		* ^________________________________________________________________________________________ Beatmap MD5 mde^______mods
		* 52 71 04 00
		^ ^_ BeatmapID
		*/
		pack.Unmarshal(&s.User.Status.Status, &s.User.Status.Text, &s.User.Status.MapMD5)
	case pid.OsuSendIRCMessage:
		HandleMessage(pack, s)
	case pid.OsuChannelJoin:
		// Just a string containing the channel name to join.
	case pid.OsuExit:
		UserQuit(s)
		delAfter = true
	}
	return
}
