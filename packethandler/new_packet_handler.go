package packethandler

import (
	"github.com/bnch/bancho/inbound"
)

// Packet IDs done by the osu! client.
const (
	PacketChangeStatus = 0
	PacketChatMessage  = 1
	PacketDisconnect   = 2
	// What's the difference between this and PacketBotnetAvailableBeatmaps?
	// Well, to be honest, I have no idea.
	// I tested the difference between the values. I got a difference of two maps.
	// So perhaps packet 63 does not have certain kind of maps?
	PacketBeatmapBotnet           = 3
	PacketRequestUpdates          = 4
	PacketRequestSpectator        = 16
	PacketPrivateMessage          = 25
	PacketJoinLobby               = 30
	PacketRequestChannelJoin      = 63
	PacketBotnetAvailableBeatmaps = 68
	PacketAddToFriends            = 73
	PacketRemoveFromFriends       = 74
	PacketPartChannel             = 78
	PacketRequestOnlineUsersData  = 85
)

// NewPacketHandler handles inbound packets, and sends them to be analysed by functions, or if it's short by a case itself.
func NewPacketHandler(pack inbound.BasePacket, s *Session) {
	switch pack.ID {
	case PacketChangeStatus:
		/* E.G.:
		* 00 00 00 55 00 00 00 06 0b 27 6f 73 75 21 20 70 6c 61 79 20 4f 6d 6f 69 20 2d 20 53 6e 6f 77 20 44 72 69 76 65 20 5b 53 6e 6f 77 44 72 69 76 65 5d
		* ^____________ header ^  ^_____________________________________________________________________________________________________________ Action text
		*                      |_ Status (common.Status constants)
		* 0b 20 36 34 39 32 30 34 33 61 33 31 33 61 37 34 30 61 61 36 38 32 66 63 61 63 64 38 64 66 36 38 38 31 00 08 00 00 00 52 71
		* ^_____________________________________________________________________________________________________________ Beatmap MD5
		* 04 00
		^ ^_ IDK
		*/
		pack.Unmarshal(&s.User.Status.Status, &s.User.Status.Text, &s.User.Status.MapMD5)
	case PacketChatMessage:
		/* E.G.:
		* 01 00 00 16 00 00 00 0b 00 0b 08 64 69 6f 20 63 6e 61 65 0b 04 23 6f 73 75 00 00 00 00
		* ^____________ header
		 */
	case PacketRequestChannelJoin:
		// Just a string containing the channel name to join.
	}
}
