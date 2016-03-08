package packethandler

import (
	"sync"

	"github.com/bnch/bancho/packets"
)

// Warning: very low-level code which you probably don't need to touch ahead.

var streams map[string]*Stream
var streamsMutex *sync.Mutex

// Stream is a way to handle sending of packets to multiple users.
type Stream struct {
	name        string
	subscribers []string
	subsMutex   *sync.Mutex
}

// NewStream creates a new default stream
func NewStream(name string) *Stream {
	s := &Stream{
		name:      name,
		subsMutex: &sync.Mutex{},
	}
	streamsMutex.Lock()
	streams[name] = s
	streamsMutex.Unlock()
	return s
}

// GetStream returns an existing stream if it does exist, nil otherwise.
func GetStream(name string) *Stream {
	streamsMutex.Lock()
	defer streamsMutex.Unlock()
	if stream, ok := streams[name]; ok {
		return stream
	}
	return nil
}

// GetInitialisedStream returns a valid stream in all cases.
// If there's no stream with such name, it creates it.
func GetInitialisedStream(name string) *Stream {
	s := GetStream(name)
	if s == nil {
		s = NewStream(name)
	}
	return s
}

// Delete erases the stream.
func (s *Stream) Delete() {
	go s.delete()
}
func (s *Stream) delete() {
	streamsMutex.Lock()
	defer streamsMutex.Unlock()
	delete(streams, s.name)
}

// Subscribe subscribes an user to a channel. Here an user is its token.
func (s *Stream) Subscribe(u string) {
	go s.subscribe(u)
}
func (s *Stream) subscribe(u string) {
	s.subsMutex.Lock()
	defer s.subsMutex.Unlock()
	if !s.isSubscribed(u) {
		s.subscribers = append(s.subscribers, u)
	}
}

// Unsubscribe removes an user from the stream.
func (s *Stream) Unsubscribe(u string) {
	go s.unsubscribe(u)
}
func (s *Stream) unsubscribe(u string) {
	s.subsMutex.Lock()
	defer s.subsMutex.Unlock()
	for i, subscriber := range s.subscribers {
		if subscriber == u {
			s.subscribers = append(s.subscribers[:i], s.subscribers[i+1:]...)
			break
		}
	}
}

// Subscribers is a function because we want to make it sure to be read-only.
func (s *Stream) Subscribers() []string {
	s.subsMutex.Lock()
	defer s.subsMutex.Unlock()
	return s.subscribers
}

// IsSubscribed checks whether an user is already subscribed.
func (s *Stream) IsSubscribed(u string) bool {
	s.subsMutex.Lock()
	defer s.subsMutex.Unlock()
	return s.isSubscribed(u)
}

func (s *Stream) isSubscribed(u string) bool {
	for _, v := range s.subscribers {
		if u == v {
			return true
		}
	}
	return false
}

// Name returns the name of the stream.
func (s *Stream) Name() string {
	// Thanks god this doesn't have mutex memes.
	return s.name
}

// Send sends something to all the users in the stream.
func (s *Stream) Send(p packets.Packet) {
	s.send(p)
}
func (s *Stream) send(p packets.Packet) {
	lSessions := CopySessions()
	for _, u := range s.subscribers {
		sess, ok := lSessions[u]
		if !ok {
			s.Unsubscribe(u)
			continue
		}
		sess.Push(p)
	}
}
