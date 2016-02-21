package packethandler

import (
	"github.com/bnch/bancho/common"
)

// User represents an user online on bancho.
type User struct {
	ID     int32
	Name   string
	Token  string
	Status common.Status
}
