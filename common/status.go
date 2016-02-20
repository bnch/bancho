package common

// These are the various statuses an user can have.
const (
	StatusIdle = iota
	StatusAfk
	StatusPlaying
	StatusEditing
	StatusModding
	StatusMultiplayer
	StatusWatching
	StatusUnknown
	StatusTesting
	StatusSubmitting
	StatusPaused
	StatusLobby
	StatusMultiplaying
	StatusOsuDirect
)

// Status indicates the action currently being done by the user.
type Status struct {
	Status byte
	Text   string
	MapMD5 string
}
