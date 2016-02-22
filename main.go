package main

import (
	"fmt"
	"github.com/bnch/bancho/cmd"
	"github.com/bnch/bancho/conf"
	"github.com/bnch/bancho/models"
	"github.com/bnch/bancho/web"
	"github.com/codegangsta/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "bancho"
	app.Usage = "custom bancho implementation"
	app.Action = func(c *cli.Context) {
		if !checkConf() {
			return
		}
		fmt.Println("== Welcome to bancho. ==")
		db, err := models.CreateDB()
		if err != nil {
			panic(err)
		}
		models.Migrate(db)
		web.Start(":3000", ":10443")
	}
	app.Commands = []cli.Command{
		{
			Name:   "mkuser",
			Action: cmd.MkUser,
		},
	}
	app.Run(os.Args)
}

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
