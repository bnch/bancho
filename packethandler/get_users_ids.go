package packethandler

// GetUserIDs retrieves all the online users' IDs.
func GetUserIDs() []int32 {
	var users []int32
	for _, sess := range Sessions {
		users = append(users, sess.User.ID)
	}
	return users
}
