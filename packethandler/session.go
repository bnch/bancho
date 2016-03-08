package packethandler

import (
	"bytes"
	"fmt"
	"os"
	"time"

	"github.com/bnch/bancho/packets"
	"github.com/bnch/banchoreader/lib"
)

// Session is an alive connection of a logged in user.
type Session struct {
	stream      *bytes.Buffer
	User        User
	LastRequest time.Time
}

// Push appends an element to the current session.
func (s Session) Push(val ...packets.Packet) {
	if s.stream == nil {
		s.stream = new(bytes.Buffer)
	}
	dumper := banchoreader.New()
	dumper.Colored = true
	fmt.Printf("> To: %s\n", s.User.Name)
	for _, v := range val {
		var c bool
		for _, ignored := range v.Ignored {
			if s.User.Token == ignored {
				c = true
				break
			}
		}
		if c {
			continue
		}
		s.stream.Write(v.Content)
		dumper.Dump(os.Stdout, v.Content)
	}
}

// NewSession generates a new session.
func NewSession(u User) (*Session, string) {
	var tok string
	for {
		tok = GenerateGUID()
		// Make sure token does not already exist
		if _, ok := sessions[tok]; !ok {
			break
		}
	}
	u.Token = tok
	return &Session{
		stream:      new(bytes.Buffer),
		User:        u,
		LastRequest: time.Now(),
	}, u.Token
}

// Sessions is a map of connections to the server via the bancho protocol.
var sessions map[string]*Session
