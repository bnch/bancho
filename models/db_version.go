package models

// DBVer is a struct which only scope is to help manage migrations.
type DBVer struct {
	ID      int
	Version uint64 // as it's a UNIX timestamp, we need to take something high.
}
