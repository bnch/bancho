package packethandler

import (
	"github.com/bnch/bancho/packets"
)

// UserQuit tells everyone that an user has quit.
func UserQuit(ps packSess) {
	st := GetStream("all")
	st.Unsubscribe(ps.s.User.Token)

	// Don't tell other users we quit if there's still someone with our identity online.
	count := 0
	myID := ps.s.User.ID
	for _, session := range CopySessions() {
		if session.User.ID == myID {
			count++
		}
	}
	if count > 1 {
		return
	}

	st.Send(packets.UserQuit(ps.s.User.ID))
}

// DeleteCompletely deletes all traces of the user in the RAM.
func DeleteCompletely(token string) {
	sessionsMutex.RLock()
	uid := sessions[token].User.ID
	sessionsMutex.RUnlock()

	sessionsMutex.Lock()
	delete(sessions, token)
	sessionsMutex.Unlock()

	uidToSessionMutex.Lock()
	delete(uidToSession, uid)
	uidToSessionMutex.Unlock()
}
