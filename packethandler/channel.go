package packethandler

import (
	"github.com/bnch/bancho/inbound"
	"github.com/bnch/bancho/models"
	"github.com/bnch/bancho/packets"
)

// HandleChannelJoin handles requests to join a channel.
func HandleChannelJoin(pack inbound.BasePacket, s *Session) {
	var channelToJoin string
	pack.Unmarshal(&channelToJoin)

	var chans []models.Channel
	db.Find(&chans)
	var found bool
	for _, val := range chans {
		if val.Name == channelToJoin {
			found = true
			break
		}
	}
	if found {
		st := GetInitialisedStream("chan/" + channelToJoin)
		st.Subscribe(s.User.Token)
		s.Push(packets.ChannelJoin(channelToJoin))
	} else {
		SendMessage(s.User.Token, "No such channel exists!")
		s.Push(packets.ChannelRemove(channelToJoin))
	}
}
