package packethandler

import (
	"github.com/bnch/bancho/common"
	"github.com/bnch/bancho/models"
	"github.com/bnch/bancho/packethandler/logindata"
	"github.com/bnch/bancho/packets"
)

const protocolVersion = 19

// Login logs the user into bancho. Returns the osu! token and any eventual error.
func Login(l logindata.LoginData) (string, bool, error) {
	sess, guid := NewSession(User{})

	user := models.User{}
	db.Where(&models.User{
		Username: l.Username,
	}).First(&user)

	if !common.IsSamePass(l.Password, user.Password) {
		sess.Push(
			packets.UserID(packets.LoginFailed),
		)
		return guid, true, nil
	}
	if (user.Permissions & models.PermissionBanned) != 0 {
		sess.Push(
			packets.UserID(packets.LoginBanned),
		)
		return guid, true, nil
	}

	privileges := uint32(packets.PrivilegeSupporter)
	if user.IsAdmin() || user.IsModerator() {
		privileges = packets.PrivilegeGMTSupporter
	}

	banchoUser := &sess.User
	banchoUser.Colour = user.GetColour()
	banchoUser.Country = 108
	banchoUser.UTCOffset = 24
	banchoUser.ID = int32(user.ID)
	banchoUser.Rank = 1337
	banchoUser.Name = user.Username

	sess.Push(
		packets.SilenceClient(0),
		packets.UserID(banchoUser.ID),
		packets.ChoProtocol(protocolVersion),
		packets.UserPrivileges(privileges),
		packets.FriendList(GetUserFriends(user.ID)),
		packets.UserData(banchoUser.ToUserDataInfo()),
		packets.UserDataFull(packets.UserDataFullInfo{
			ID:         int32(user.ID),
			Action:     common.StatusIdle,
			Mods:       0,
			GameMode:   packets.ModeStandard,
			Score:      147200000,
			Accuracy:   13.37,
			Playcount:  1231,
			TotalScore: 1200200000,
			Rank:       1337,
			PP:         0, // 0 because not implemented
		}),
		packets.OnlinePlayers(GetUserIDs()),
		packets.ChannelJoin("#osu"),
		packets.ChannelJoin("#announce"),
	)

	channelsMutex.RLock()
	for _, c := range channels {
		st := GetInitialisedStream("chan/" + c.Name)
		subs := len(st.Subscribers())
		var subsUint uint16
		if subs > 65535 {
			subsUint = 9001
		} else {
			subsUint = uint16(subs)
		}
		sess.Push(
			packets.ChannelTitle(c.Name, c.Description, subsUint),
		)
	}
	channelsMutex.RUnlock()
	sess.Push(packets.ChannelListingComplete())

	uidToSessionMutex.Lock()
	uidToSession[int32(user.ID)] = sess
	uidToSessionMutex.Unlock()

	s := GetStream("all")
	s.Subscribe(guid)
	go sendUserPresence(s, int32(user.ID))

	GetStream("chan/#osu").Subscribe(guid)
	GetStream("chan/#announce").Subscribe(guid)

	return guid, false, nil
}

func sendUserPresence(s *Stream, uid int32) {
	count := 0
	for _, session := range CopySessions() {
		if session.User.ID == uid {
			count++
		}
	}
	if count < 2 {
		s.Send(packets.UserPresence(uid))
	}
}
