package models

import (
	"github.com/jinzhu/gorm"
)

var migrations = map[uint64]func(gorm.DB){
	1455998503: func(db gorm.DB) {
		db.CreateTable(&User{})
	},
	1456178412: func(db gorm.DB) {
		// Users start from ID 3.
		// This is because if an user comes to have ID == 2, then they become peppy and thus can't be contacted
		// (any attempt to message them will result in a browser opening http://osu.ppy.sh/p/doyoureallywanttoaskpeppy)
		db.Exec("ALTER TABLE users AUTO_INCREMENT=3;")
	},
	1456756895: func(db gorm.DB) {
		db.CreateTable(&UserFriendship{})
	},
	1457036659: func(db gorm.DB) {
		db.CreateTable(&Channel{})
		db.Create(&Channel{
			Name:        "#osu",
			Description: "Main channel for discussion about anything and everything.",
		})
		db.Create(&Channel{
			Name:        "#announce",
			Description: "The channnel where the announcements should appear. Not really.",
		})
	},
	1457125054: func(db gorm.DB) {
		db.CreateTable(&UserStats{})
		users := []User{}
		db.Find(&users)
		for _, u := range users {
			db.Create(&UserStats{
				ID: u.ID,
			})
		}
	},
	1457180378: func(db gorm.DB) {
		db.CreateTable(&Leaderboard{})
		BuildLeaderboard(db)
	},
}
