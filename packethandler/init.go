package packethandler

import (
	"sync"

	"github.com/bnch/bancho/models"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

// SetUp is needed to make things work and should always be called unless you have a very good fucking reason for not doing it.
func SetUp() {
	sessionsMutex = &sync.Mutex{}
	streamsMutex = &sync.Mutex{}
	sessions = make(map[string]*Session)
	streams = make(map[string]*Stream)
	var err error
	db, err = models.CreateDB()
	if err != nil {
		panic(err)
	}
	NewStream("all")
	go Prune()
}
