package packethandler

import (
	"bytes"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/bnch/bancho/packets"
	"github.com/bnch/banchoreader/lib"
)

// Session is an alive connection of a logged in user.
type Session struct {
	stream      *bytes.Buffer
	User        User
	LastRequest time.Time
	Mutex       *sync.Mutex
}

// Push appends an element to the current session.
func (s Session) Push(val ...packets.Packet) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()
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
	sess := &Session{
		stream:      new(bytes.Buffer),
		User:        u,
		LastRequest: time.Now(),
		Mutex:       &sync.Mutex{},
	}
	sessionsMutex.Lock()
	defer sessionsMutex.Unlock()
	sessions[tok] = sess
	return sess, tok
}

// GetSession retrieves a session from the available ones.
func GetSession(sessName string) *Session {
	sessionsMutex.Lock()
	defer sessionsMutex.Unlock()
	return sessions[sessName]
}

// CopySessions can be used to get an independent copy of sessions, without need to use the sessionMutex to modify it.
func CopySessions() map[string]*Session {
	sessionsMutex.Lock()
	defer sessionsMutex.Unlock()
	ret := make(map[string]*Session, len(sessions))
	for k, v := range sessions {
		ret[k] = v
	}
	return ret
}

// Sessions is a map of connections to the server via the bancho protocol.
var sessions map[string]*Session
var sessionsMutex *sync.Mutex
