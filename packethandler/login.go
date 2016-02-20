package packethandler

import (
	"git.zxq.co/ripple/go-bancho/common"
	"git.zxq.co/ripple/go-bancho/packethandler/logindata"
	"git.zxq.co/ripple/go-bancho/packets"
)

const protocolVersion = 19

// Login logs the user into bancho. Returns the osu! token and any eventual error.
func Login(l logindata.LoginData) (string, error) {
	sess, guid := NewSession(User{
		ID:   1337,
		Name: "Howl",
	})
	Sessions[guid] = sess
	Sessions[guid].Push(
		packets.SilenceClient(0),
		packets.UserID(1),
		packets.ChoProtocol(protocolVersion),
		packets.UserPrivileges(packets.PrivilegeGMTSupporter),
		packets.FriendList([]int32{2}),
		packets.UserData(packets.UserDataInfo{
			ID:         1,
			PlayerName: "Howl",
			UTCOffset:  25,
			Country:    108,
			Colour:     packets.ColourAdmin,
			Longitude:  0,
			Latitude:   0,
			Rank:       9,
		}),
		packets.UserDataFull(packets.UserDataFullInfo{
			ID:         1,
			Action:     common.StatusIdle,
			Mods:       0,
			GameMode:   packets.ModeStandard,
			Score:      147200000,
			Accuracy:   13.37,
			Playcount:  1231,
			TotalScore: 1200200000,
			Rank:       9,
			PP:         0, // 0 because not implemented
		}),
		packets.UserData(packets.UserDataInfo{
			ID:         2,
			PlayerName: "Nyo",
			UTCOffset:  25,
			Country:    108,
			Colour:     packets.ColourNormal,
			Longitude:  0,
			Latitude:   0,
			Rank:       8,
		}),
		packets.OnlinePlayers([]int32{
			1,
			2,
		}),
		packets.P89(),
		packets.ChannelJoin("#osu"),
		packets.ChannelJoin("#announce"),
		packets.ChannelTitle("#osu", "WELCOME TO THE DANK MEMES", 2),
		packets.ChannelTitle("#announce", "WELCOME TO THE DANK MEMES, PART 2", 1337),
		packets.ChannelTitle("#puckfeppy", "Ayy Lmao", 1338),
	)
	return guid, nil
}
