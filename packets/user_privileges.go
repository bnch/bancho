package packets

import (
	"github.com/bnch/bancho/pid"
)

// User ranks
const (
	PrivilegeNormal       = 1
	PrivilegeGMT          = 2
	PrivilegeSupporter    = 4
	PrivilegeGMTSupporter = PrivilegeSupporter | PrivilegeGMT
)

// UserPrivileges returns a packet with the privileges of the user. Privileges can be picked using the related constants.
func UserPrivileges(privileges uint32) Packet {
	return MakePacket(pid.BanchoLoginPermissions, 4, privileges)
}
