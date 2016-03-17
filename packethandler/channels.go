package packethandler

import (
	"sync"

	"github.com/bnch/bancho/models"
)

var channels []models.Channel
var channelsMutex *sync.RWMutex
