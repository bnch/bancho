package avatarserver

import (
	"io/ioutil"
	"os"
)

// SetUp makes sure the folders where to write the avatars exist.
func SetUp() error {
	err := os.MkdirAll("data/avatars", 0755)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile("data/avatars/default.png", MustAsset("data/default.png"), 0644)
	if err != nil {
		return err
	}
	return nil
}
