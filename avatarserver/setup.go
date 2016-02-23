package avatarserver

import (
    "os"
)

// SetUp makes sure the folders where to write the avatars exist.
func SetUp() error {
    err := os.MkdirAll("data/avatars", 0755)
	if err != nil {
		return err
	}
	return nil
}