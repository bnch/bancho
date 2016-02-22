package main

import (
	"github.com/bnch/bancho/cmd"
	"github.com/codegangsta/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "bancho"
	app.Usage = "custom bancho implementation"
	app.Action = cmd.Start
	app.Commands = []cli.Command{
		{
			Name:   "mkuser",
			Action: cmd.MkUser,
		},
	}
	app.Run(os.Args)
}
