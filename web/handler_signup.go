package web

import (
	"github.com/asaskevich/govalidator"
	"github.com/bnch/bancho/common"
	"github.com/bnch/bancho/models"
	"github.com/gin-gonic/gin"
	"regexp"
	"strings"
)

const usernameRegexString = "^[a-zA-Z0-9_\\[\\] -]+$"

func signupHandler(c *gin.Context) {
	user := c.PostForm("username")
	pass := c.PostForm("password")
	pass2 := c.PostForm("password2")
	email := c.PostForm("email")

	// The most boilerplate, yet necessary, code in all software...
	// 1. Check that the username is valid.
	usernameRegex := regexp.MustCompile(usernameRegexString)
	if !usernameRegex.Match([]byte(user)) {
		c.String(200, "Username can only contain alphanumeric characters, low dashes, dashes and brackets."+user)
		return
	}

	// 2. Check that the passwords are the same.
	if pass != pass2 {
		c.String(200, "Passwords are not the same!")
		return
	}

	// 3. Check password to be long enough. We ain't cunts, so we don't check
	// the user has written special characters in the password and blah blah blah.
	if len(pass) < 8 {
		c.String(200, "Did you know that short passwords are the most likely to be cracked? Please use a password at least 6 characters long.")
		return
	}

	// 4. Check email to be valid
	if !govalidator.IsEmail(email) {
		c.String(200, "We aren't so dumb to let THAT pass through as an email. Please invent a better one. Bonus if we can send mails to it.")
		return
	}

	// 5. Check for an user with the same username
	u := models.User{}
	db.Where("username = ? or email = ?", user, pass).First(&u)
	if !db.NewRecord(u) {
		var samething string
		if strings.ToLower(user) == strings.ToLower(u.Username) {
			samething = "username"
		} else {
			samething = "email"
		}
		c.String(200, "An user with that same "+samething+" already exists!")
		return
	}

	pass = common.CryptPass(pass)
	db.Create(&models.User{
		Username:    user,
		Password:    pass,
		Permissions: models.PermissionAdmin,
	})
	c.String(200, "success")
}
