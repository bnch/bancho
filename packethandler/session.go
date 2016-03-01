package packethandler

import (
	"github.com/bnch/bancho/packets"
	"time"
)

// Session is an alive connection of a logged in user.
type Session struct {
	stream      *packetCollection
	User        User
	LastRequest time.Time
}

// Push appends an element to the current session.
func (s Session) Push(val ...packets.Packet) {
	if s.stream == nil {
		s.stream = &packetCollection{}
	}
	s.stream.Push(val...)
}

// NewSession generates a new session.
func NewSession(u User) (*Session, string) {
	var tok string
	for {
		tok = GenerateGUID()
		// Make sure token does not already exist
		if _, ok := Sessions[tok]; !ok {
			break
		}
	}
	u.Token = tok
	return &Session{
		stream:      new(packetCollection),
		User:        u,
		LastRequest: time.Now(),
	}, u.Token
}

// Sessions is a map of connections to the server via the bancho protocol.
var Sessions map[string]*Session
