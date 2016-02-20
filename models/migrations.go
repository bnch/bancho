package models

import (
	"github.com/jinzhu/gorm"
)

var migrations = map[uint64]func(gorm.DB){
	1455998503: func(db gorm.DB) {
		db.CreateTable(&User{})
	},
}
