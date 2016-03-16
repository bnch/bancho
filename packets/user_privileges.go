package packets

import (
	"github.com/bnch/bancho/pid"
)

// User ranks
const (
	PrivilegeNormalest = 0
	PrivilegeNormal    = 1 << iota
	PrivilegeGMT
	PrivilegeSupporter
	PrivilegePeppy
	PrivilegeAdmin
	PrivilegeTournamentStaff

	PrivilegeGMTSupporter = PrivilegeSupporter | PrivilegeGMT
)

// UserPrivileges returns a packet with the privileges of the user. Privileges can be picked using the related constants.
func UserPrivileges(privileges uint32) Packet {
	return MakePacket(pid.BanchoLoginPermissions, 4, privileges)
}
