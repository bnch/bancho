package packethandler

import (
	"github.com/bnch/bancho/models"
	"github.com/jinzhu/gorm"
)

var db gorm.DB

// SetUp is needed to make things work and should always be called unless you have a very good fucking reason for not doing it.
func SetUp() {
	Sessions = make(map[string]Session)
	var err error
	db, err = models.CreateDB()
	if err != nil {
		panic(err)
	}
	go Prune()
}
