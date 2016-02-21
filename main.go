package main

import (
	"fmt"
	"github.com/bnch/bancho/cmd"
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
		fmt.Println("== Welcome to bancho. ==")
		db, err := models.CreateDB()
		if err != nil {
			panic(err)
		}
		models.Migrate(db)
		web.Start(":3000", "10443")
	}
	app.Commands = []cli.Command{
		{
			Name:   "mkuser",
			Action: cmd.MkUser,
		},
	}
	app.Run(os.Args)
}
