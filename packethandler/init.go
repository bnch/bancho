package packethandler

import (
	"github.com/bnch/bancho/models"
	"github.com/jinzhu/gorm"
)

var db gorm.DB

func init() {
	Sessions = make(map[string]Session)
	var err error
	db, err = models.CreateDB()
	if err != nil {
		panic(err)
	}
}
