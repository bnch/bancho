package logindata

import (
	"errors"
	"strconv"
	"strings"
)

// LoginData is the data received by the osu! client upon a login request to bancho.
type LoginData struct {
	Username       string
	Password       string
	HardwareData   []string
	HardwareHashes []string
}

// Unmarshal creates a new LoginData with the data passed.
func Unmarshal(input []byte) (l LoginData, e error) {
	lines := strings.Split(string(input), "\n")
	if len(lines) != 4 {
		e = errors.New("logindata: cannot unmarshal, got " + strconv.Itoa(len(lines)) + " lines as an input, want 4")
		return
	}
	l.Username = lines[0]
	l.Password = lines[1]
	l.HardwareData = strings.Split(lines[2], "|")
	l.HardwareHashes = strings.Split(lines[3], ":")
	return
}
