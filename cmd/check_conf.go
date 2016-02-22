package cmd

import (
	"fmt"
	"github.com/bnch/bancho/conf"
	"os"
)

func checkConf() bool {
	if _, err := os.Stat("bancho.ini"); os.IsNotExist(err) {
		err = conf.WriteSampleConf()
		if err != nil {
			panic(err)
		}
		fmt.Println("We have made a default config file for you. Come back when you're all set up.")
		return false
	}
	return true
}
