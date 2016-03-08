package packethandler

// GetUserIDs retrieves all the online users' IDs.
func GetUserIDs() []int32 {
	users := make([]int32, len(sessions)+1)
	users[0] = BotID
	i := 1
	for _, sess := range CopySessions() {
		if sess != nil && sess.User.ID != 0 {
			if i >= len(users) {
				users = append(users, sess.User.ID)
			} else {
				users[i] = sess.User.ID
			}
			i++
		}
	}
	return users[:i]
}
