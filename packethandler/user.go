package packethandler

import (
	"git.zxq.co/ripple/go-bancho/common"
)

// User represents an user online on bancho.
type User struct {
	ID     int32
	Name   string
	Token  string
	Status common.Status
}
