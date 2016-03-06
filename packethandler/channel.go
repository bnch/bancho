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

// HandleChannelPart handles requests to part from a channel.
func HandleChannelPart(pack inbound.BasePacket, s *Session) {
	var channelToPart string
	pack.Unmarshal(&channelToPart)

	// I'll confess: I made these just so I could chuckle a bit.
	st := GetStream("chan/" + channelToPart)
	if st == nil {
		SendMessage(s.User.Token, "w00t p00t! That channel doesn't even exist, yet you're requesting to part from it? What are you, an akerino?!")
	} else if !st.IsSubscribed(s.User.Token) {
		SendMessage(s.User.Token, "dude, the fuck? You aren't even in that channel, how am I supposed to make you part from it?")
	} else {
		st.Unsubscribe(s.User.Token)
	}
}
