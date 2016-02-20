package cmd

import (
	"git.zxq.co/ripple/go-bancho/common"
	"git.zxq.co/ripple/go-bancho/models"
	"github.com/codegangsta/cli"
	"github.com/jinzhu/gorm"
)

// MkUser puts an user in the database.
func MkUser(c *cli.Context) {
	a := c.Args()
	user := a[0]
	pass := a[1]
	var err error
	var db gorm.DB
	db, err = models.CreateDB()
	if err != nil {
		panic(err)
	}
	pass = common.CryptPass(pass)
	db.Create(&models.User{
		Username:    user,
		Password:    pass,
		Permissions: models.PermissionAdmin,
	})
}
