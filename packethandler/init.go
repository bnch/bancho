package packethandler

import (
	"sync"

	"github.com/bnch/bancho/models"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

// SetUp is needed to make things work and should always be called unless you have a very good fucking reason for not doing it.
func SetUp() {

	sessionsMutex = &sync.RWMutex{}
	streamsMutex = &sync.RWMutex{}
	uidToSessionMutex = &sync.RWMutex{}
	channelsMutex = &sync.RWMutex{}

	sessions = make(map[string]*Session)
	streams = make(map[string]*Stream)
	uidToSession = make(map[int32]*Session)

	var err error
	db, err = models.CreateDB()
	if err != nil {
		panic(err)
	}

	db.Find(&channels)
	for _, c := range channels {
		GetInitialisedStream("chan/" + c.Name)
	}

	NewStream("all")
	go Prune()
}
