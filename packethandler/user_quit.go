package packethandler

import (
	"github.com/bnch/bancho/packets"
)

// UserQuit tells everyone that an user has quit.
func UserQuit(s *Session) {
	st := GetStream("all")
	st.Unsubscribe(s.User.Token)
	st.Send(packets.UserQuit(s.User.ID))
}
