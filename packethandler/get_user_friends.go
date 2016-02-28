package packethandler

import (
	"github.com/bnch/bancho/models"
)

// GetUserFriends returns an int32 array containing all the friends of an user, be them online or not.
func GetUserFriends(userID int) []int32 {
	var lovers []models.UserFriendship
	db.Where(&models.UserFriendship{
		Lover: userID,
	}).Find(&lovers)

	var endArr []int32
	for _, lover := range lovers {
		endArr = append(endArr, int32(lover.Lover))
	}

	return endArr
}
