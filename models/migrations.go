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
}
