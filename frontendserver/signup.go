package frontendserver

import (
	"regexp"
	"strings"

	"github.com/asaskevich/govalidator"
	"github.com/bnch/bancho/common"
	"github.com/bnch/bancho/models"
	"github.com/gin-gonic/gin"
)

const usernameRegexString = "^[a-zA-Z0-9_\\[\\] -]{1,20}$"

func signupPOST(c *gin.Context) {
	user := c.PostForm("username")
	pass := c.PostForm("password")
	pass2 := c.PostForm("password2")
	email := c.PostForm("email")

	// The most boilerplate, yet necessary, code in all software...
	// 1. Check that the username is valid.
	usernameRegex := regexp.MustCompile(usernameRegexString)
	if !usernameRegex.Match([]byte(user)) {
		signupError(c, "Username can only contain alphanumeric characters, low dashes, dashes and brackets.")
		return
	}

	// 2. Check that the passwords are the same.
	if pass != pass2 {
		signupError(c, "Passwords are not the same!")
		return
	}

	// 3. Check password to be long enough. We ain't cunts, so we don't check
	// the user has written special characters in the password and blah blah blah.
	if len(pass) < 8 {
		signupError(c, "Did you know that short passwords are the most likely to be cracked? Please use a password at least 8 characters long.")
		return
	}

	// 4. Check email to be valid
	if !govalidator.IsEmail(email) {
		signupError(c, "We aren't so dumb to let THAT pass through as an email. Please invent a better one. Bonus if we can send mails to it.")
		return
	}

	// 5. Check for an user with the same username
	u := models.User{}
	db.Where("username = ? or email = ?", user, email).First(&u)
	if !db.NewRecord(u) {
		var samething string
		if strings.ToLower(user) == strings.ToLower(u.Username) {
			samething = "username"
		} else {
			samething = "email"
		}
		signupError(c, "An user with that same "+samething+" already exists!")
		return
	}

	pass = common.CryptPass(pass)
	dbUser := models.User{
		Username:    user,
		Password:    pass,
		Permissions: 0,
	}
	db.Create(&dbUser)
	go func() {
		dbUserStats := models.UserStats{
			ID: dbUser.ID,
		}
		db.Create(&dbUserStats)
		dbUserStats.UpdateLeaderboard(0, db)
		dbUserStats.UpdateLeaderboard(1, db)
		dbUserStats.UpdateLeaderboard(2, db)
		dbUserStats.UpdateLeaderboard(3, db)
	}()
	serveTemplate("signup", gin.H{
		"Title": "Sign up",
		"Status": gin.H{
			"Success": "Successfully signed you up! You should now be able to login.",
		},
	}, 200, c)
}
func signupError(c *gin.Context, failMessage string) {
	serveTemplate("signup", gin.H{
		"Title": "Sign up",
		"Status": gin.H{
			"Failure": failMessage,
		},
		"PostData": gin.H{
			"Username":  c.PostForm("username"),
			"Password":  c.PostForm("password"),
			"Password2": c.PostForm("password2"),
			"Email":     c.PostForm("email"),
		},
	}, 200, c)
}
func signupGET(c *gin.Context) {
	serveTemplate("signup", gin.H{
		"Title": "Sign up",
	}, 200, c)
}
