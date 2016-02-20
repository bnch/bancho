package main

import (
	"fmt"
	"git.zxq.co/ripple/go-bancho/cmd"
	"git.zxq.co/ripple/go-bancho/models"
	"git.zxq.co/ripple/go-bancho/web"
	"github.com/codegangsta/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "go-bancho"
	app.Usage = "custom bancho implementation"
	app.Action = func(c *cli.Context) {
		fmt.Println("== Welcome to go-bancho. ==")
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
