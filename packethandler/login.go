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

	Sessions[guid].Push(
		packets.SilenceClient(0),
		packets.UserID(int32(user.ID)),
		packets.ChoProtocol(protocolVersion),
		packets.UserPrivileges(packets.PrivilegeGMTSupporter),
		packets.FriendList([]int32{9001}),
		packets.UserData(packets.UserDataInfo{
			ID:         int32(user.ID),
			PlayerName: user.Username,
			UTCOffset:  25,
			Country:    108,
			Colour:     packets.ColourAdmin,
			Longitude:  0,
			Latitude:   0,
			Rank:       1337,
		}),
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
		packets.UserData(packets.UserDataInfo{
			ID:         9001,
			PlayerName: "Michele Satori",
			UTCOffset:  25,
			Country:    108,
			Colour:     packets.ColourNormal,
			Longitude:  0,
			Latitude:   0,
			Rank:       8,
		}),
		packets.OnlinePlayers([]int32{
			int32(user.ID),
			9001,
		}),
		packets.ChannelListingComplete(),
		packets.ChannelJoin("#osu"),
		packets.ChannelJoin("#announce"),
		packets.ChannelTitle("#osu", "WELCOME TO THE DANK MEMES", 2),
		packets.ChannelTitle("#announce", "WELCOME TO THE DANK MEMES, PART 2", 1337),
		packets.ChannelTitle("#puckfeppy", "Ayy Lmao", 1338),
	)
	return guid, nil
}
