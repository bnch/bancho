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
	pack := packets.ChatMessage(c.From, c.To, c.Content, c.UserID)
	pack.Ignored = append(pack.Ignored, s.User.Token)
	return pack
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
	st.Send(m.ToPacket(s))
}
