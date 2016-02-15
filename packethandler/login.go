package packethandler

import (
	"git.zxq.co/ripple/go-bancho/packethandler/logindata"
	"git.zxq.co/ripple/go-bancho/packets"
)

// Login logs the user into bancho. Returns the osu! token and any eventual error.
func Login(l logindata.LoginData) (string, error) {
	guid := GenerateGUID()
	Tokens[guid] = new(packetCollection)
	Tokens[guid].Push(packets.UserID(2))
	return guid, nil
}
