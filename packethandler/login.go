package packethandler

import (
	"github.com/bnch/bancho/common"
	"github.com/bnch/bancho/models"
	"github.com/bnch/bancho/packethandler/logindata"
	"github.com/bnch/bancho/packets"
)

const protocolVersion = 19

// Login logs the user into bancho. Returns the osu! token and any eventual error.
func Login(l logindata.LoginData, output *[]byte) (string, error) {
	sess, guid := NewSession(User{})
	Sessions[guid] = sess

	user := models.User{}
	db.Where(&models.User{
		Username: l.Username,
	}).First(&user)

	if !common.IsSamePass(l.Password, user.Password) {
		Sessions[guid].Push(
			packets.UserID(packets.LoginFailed),
		)
		return guid, nil
	}
	if (user.Permissions & models.PermissionBanned) != 0 {
		Sessions[guid].Push(
			packets.UserID(packets.LoginBanned),
		)
		return guid, nil
	}
	
	privileges := uint32(packets.PrivilegeSupporter)
	if user.IsAdmin() || user.IsModerator() {
		privileges = packets.PrivilegeGMTSupporter
	}
	
	banchoUser := &Sessions[guid].User
	banchoUser.Colour = user.GetColour()
	banchoUser.Country = 108
	banchoUser.UTCOffset = 24
	banchoUser.ID = int32(user.ID)
	banchoUser.Rank = 130
	banchoUser.Username = user.Username
	
	Sessions[guid].Push(
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
		packets.ChannelListingComplete(),
		packets.ChannelJoin("#osu"),
		packets.ChannelJoin("#announce"),
		packets.ChannelTitle("#osu", "WELCOME TO THE DANK MEMES", 2),
		packets.ChannelTitle("#announce", "WELCOME TO THE DANK MEMES, PART 2", 1337),
		packets.ChannelTitle("#puckfeppy", "Ayy Lmao", 1338),
	)
	
	Broadcast(packets.UserPresence(int32(user.ID)), guid)
	
	return guid, nil
}
