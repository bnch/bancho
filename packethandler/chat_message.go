package packethandler

import (
	"github.com/bnch/bancho/inbound"
	"github.com/bnch/bancho/packets"
)

// ChatMessage is a message in the osu! chat.
type ChatMessage struct {
	From    string
	To      string
	Content string
	UserID  int32
}

// ToPacket converts a ChatMessage to a packets.Packet.
func (c ChatMessage) ToPacket(s *Session) packets.Packet {
	pack := c.ToPacketNoIgnore()
	pack.Ignored = append(pack.Ignored, s.User.Token)
	return pack
}

// ToPacketNoIgnore returns a chat message without the ignored sender. Used for PMs.
func (c ChatMessage) ToPacketNoIgnore() packets.Packet {
	return packets.ChatMessage(c.From, c.To, c.Content, c.UserID)
}

// HandleMessage broadcasts a received message to all users.
func HandleMessage(p inbound.BasePacket, s *Session) {
	m := ChatMessage{}
	p.Unmarshal(&m.From, &m.Content, &m.To, &m.UserID)

	m.From = s.User.Name
	m.UserID = s.User.ID

	st := GetStream("chan/" + m.To)
	if st == nil {
		return
	}
	if !st.IsSubscribed(s.User.Token) {
		SendMessage(s.User.Token, "You haven't joined that channel.")
	}
	st.Send(m.ToPacket(s))
}
