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
	sessions[guid] = sess

	user := models.User{}
	db.Where(&models.User{
		Username: l.Username,
	}).First(&user)

	if !common.IsSamePass(l.Password, user.Password) {
		sessions[guid].Push(
			packets.UserID(packets.LoginFailed),
		)
		return guid, true, nil
	}
	if (user.Permissions & models.PermissionBanned) != 0 {
		sessions[guid].Push(
			packets.UserID(packets.LoginBanned),
		)
		return guid, true, nil
	}

	privileges := uint32(packets.PrivilegeSupporter)
	if user.IsAdmin() || user.IsModerator() {
		privileges = packets.PrivilegeGMTSupporter
	}

	banchoUser := &sessions[guid].User
	banchoUser.Colour = user.GetColour()
	banchoUser.Country = 108
	banchoUser.UTCOffset = 24
	banchoUser.ID = int32(user.ID)
	banchoUser.Rank = 130
	banchoUser.Name = user.Username

	sessions[guid].Push(
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
		packets.ChannelTitle("#osu", "WELCOME TO THE DANK MEMES", 2),
		packets.ChannelTitle("#announce", "WELCOME TO THE DANK MEMES, PART 2", 1337),
	)

	var chans []models.Channel
	db.Find(&chans)
	for _, c := range chans {
		st := GetInitialisedStream("chan/" + c.Name)
		subs := len(st.Subscribers())
		var subsUint uint16
		if subs > 65535 {
			subsUint = 9001
		} else {
			subsUint = uint16(subs)
		}
		sessions[guid].Push(
			packets.ChannelTitle(c.Name, c.Description, subsUint),
		)
	}
	sessions[guid].Push(packets.ChannelListingComplete())

	s := GetStream("all")
	s.Subscribe(guid)
	s.Send(packets.UserPresence(int32(user.ID)))

	GetStream("chan/#osu").Subscribe(guid)
	GetStream("chan/#announce").Subscribe(guid)

	return guid, false, nil
}
