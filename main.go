package main

import (
	"fmt"
	"git.zxq.co/ripple/go-bancho/models"
	"git.zxq.co/ripple/go-bancho/web"
)

func main() {
	fmt.Println("== Welcome to go-bancho. ==")
	db, err := models.CreateDB()
	if err != nil {
		panic(err)
	}
	models.Migrate(db)
	web.Start(":3000", "10443")
}
