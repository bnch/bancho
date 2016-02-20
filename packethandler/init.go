package packethandler

import (
	"git.zxq.co/ripple/go-bancho/models"
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
