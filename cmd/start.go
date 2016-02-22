package cmd

import (
	"fmt"
	"github.com/bnch/bancho/conf"
	"github.com/bnch/bancho/models"
	"github.com/bnch/bancho/web"
	"github.com/codegangsta/cli"
)

// Start begins listening to requests.
func Start(c *cli.Context) {
	if !checkConf() {
		return
	}
	fmt.Println("== Welcome to bancho. ==")

	confFile, err := conf.Get()
	if err != nil {
		panic(err)
	}

	db, err := models.CreateDB()
	if err != nil {
		panic(err)
	}
	models.Migrate(db)

	web.Start(confFile.Ports.HTTP, confFile.Ports.HTTPS)
}
