package packethandler

import "github.com/bnch/bancho/packets"

// HandleUserStatsRequest returns to the user the stats of the users he specified in an int array.
func HandleUserStatsRequest(ps packSess) {
	var usersRequested []int32
	err := ps.p.Unmarshal(&usersRequested)
	if err != nil {
		return
	}
	for _, v := range usersRequested {
		if v == BotID {
			ps.s.Push(packets.UserData(packets.UserDataInfo{
				ID:         BotID,
				PlayerName: BotName,
			}))
			continue
		}
		u := GetSessionByID(v)
		if u != nil {
			ps.s.Push(packets.UserData(u.User.ToUserDataInfo()))
		}
	}
}
