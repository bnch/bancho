package packethandler

import (
	"github.com/bnch/bancho/packets"
)

var streams map[string]*Stream

// Stream is a way to handle sending of packets to multiple users.
type Stream struct {
	name        string
	subscribers []string
	channel     chan packets.Packet
	newUsers    chan string
	deleteUsers chan string
}

// NewStream creates a new default stream
func NewStream(name string) *Stream {
	s := &Stream{
		name:        name,
		channel:     make(chan packets.Packet),
		newUsers:    make(chan string),
		deleteUsers: make(chan string),
	}
	streams[name] = s
	go s.routine()
	return s
}

// GetStream returns an existing stream if it does exist, nil otherwise.
func GetStream(name string) *Stream {
	if stream, ok := streams[name]; ok {
		return stream
	}
	return nil
}

// Delete erases the stream.
func (s *Stream) Delete() {
	close(s.channel)
	close(s.newUsers)
	delete(streams, s.name)
}

// Subscribe subscribes an user to a channel. Here an user is its token.
func (s *Stream) Subscribe(u string) {
	s.newUsers <- u
}

// Unsubscribe removes an user from the stream.
func (s *Stream) Unsubscribe(u string) {
	s.deleteUsers <- u
}

// Subscribers is a function because we want to make it sure to be read-only.
func (s *Stream) Subscribers() []string {
	return s.subscribers
}

// Name returns the name of the stream.
func (s *Stream) Name() string {
	return s.name
}

// Send sends something to all the users in the stream.
func (s *Stream) Send(p packets.Packet) {
	s.channel <- p
}

func (s *Stream) routine() {
	for {
		// User deletion requests have the top priority.
		select {
		case x, ok := <-s.deleteUsers:
			if !ok {
				s.deleteUsers = nil
			}
			for i, subscriber := range s.subscribers {
				if subscriber == x {
					s.subscribers = append(s.subscribers[:i], s.subscribers[i+1:]...)
					break
				}
			}
		case x, ok := <-s.newUsers:
			if !ok {
				s.newUsers = nil
			}
			s.subscribers = append(s.subscribers, x)
		case x, ok := <-s.channel:
			if !ok {
				s.channel = nil
			}
			for _, u := range s.subscribers {
				sess, ok := Sessions[u]
				if !ok {
					s.Unsubscribe(u)
					continue
				}
				sess.Push(x)
			}
		}

		if s.channel == nil && s.newUsers == nil {
			break
		}
	}
}
