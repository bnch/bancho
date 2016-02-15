package packethandler

import (
	"git.zxq.co/ripple/go-bancho/packethandler/logindata"
	"git.zxq.co/ripple/go-bancho/packets"
)

const protocolVersion = 19

// Login logs the user into bancho. Returns the osu! token and any eventual error.
func Login(l logindata.LoginData) (string, error) {
	guid := GenerateGUID()
	Tokens[guid] = new(packetCollection)
	Tokens[guid].Push(
		packets.UserID(2),
		packets.SilenceClient(0),
		packets.ChoProtocol(protocolVersion),
		packets.UserPrivileges(packets.PrivilegeGMTSupporter),
		packets.FriendList([]int32{}),
	)
	return guid, nil
}
