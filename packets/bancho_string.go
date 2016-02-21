package packets

import (
	"github.com/bnch/bancho/packets/uleb128"
)

// BanchoString creates a string in the clusterfuck that is bancho.
func BanchoString(s string) []byte {
	if s == "" {
		return []byte{0}
	}
	// 11, aka 0x0b, notifies the client that what's following is a string.
	r := []byte{11}
	r = append(r, uleb128.Marshal(len(s))...)
	r = append(r, []byte(s)...)
	return r
}
