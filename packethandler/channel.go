package packethandler

import "github.com/bnch/bancho/packets"

// HandleChannelJoin handles requests to join a channel.
func HandleChannelJoin(ps packSess) {
	var channelToJoin string
	ps.p.Unmarshal(&channelToJoin)

	var found bool
	channelsMutex.RLock()
	for _, c := range channels {
		if c.Name == channelToJoin {
			found = true
			break
		}
	}
	channelsMutex.RUnlock()
	if found {
		st := GetInitialisedStream("chan/" + channelToJoin)
		st.Subscribe(ps.s.User.Token)
		ps.s.Push(packets.ChannelJoin(channelToJoin))
	} else {
		SendMessage(ps.s.User.Token, "No such channel exists!")
		ps.s.Push(packets.ChannelRemove(channelToJoin))
	}
}

// HandleChannelPart handles requests to part from a channel.
func HandleChannelPart(ps packSess) {
	var channelToPart string
	ps.p.Unmarshal(&channelToPart)

	// I'll confess: I made these just so I could chuckle a bit.
	st := GetStream("chan/" + channelToPart)
	if st == nil {
		SendMessage(ps.s.User.Token, "w00t p00t! That channel doesn't even exist, yet you're requesting to part from it? What are you, an akerino?!")
	} else if !st.IsSubscribed(ps.s.User.Token) {
		SendMessage(ps.s.User.Token, "dude, the fuck? You aren't even in that channel, how am I supposed to make you part from it?")
	} else {
		st.Unsubscribe(ps.s.User.Token)
	}
}
