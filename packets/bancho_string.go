package packets

import (
	"git.zxq.co/ripple/go-bancho/packets/uleb128"
)

// BanchoString creates a string in the clusterfuck that is bancho.
func BanchoString(s string) []byte {
	if s == "" {
		return []byte{0}
	}
	var r []byte
	r = append(r, uleb128.Marshal(len(s))...)
	r = append(r, []byte(s)...)
	return r
}
