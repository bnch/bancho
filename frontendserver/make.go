package frontendserver

import (
	"github.com/bnch/bancho/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var db gorm.DB

// Make creates a gin engine able to respond properly to requests.
func Make() *gin.Engine {
	setUpTemplates()

	var err error
	db, err = models.CreateDB()
	if err != nil {
		panic(err)
	}

	c := gin.Default()

	c.GET("/", func(c *gin.Context) {
		c.String(200, "I like memes")
	})
	c.Static("/static", "frontend/static")

	c.GET("/signup", func(c *gin.Context) {
		serveTemplate("signup", gin.H{
			"Title": "Sign up",
		}, 200, c)
	})
	c.POST("/signup", signupPOST)

	c.GET("/web/bancho_connect.php", func(c *gin.Context) {
		c.String(200, "us")
	})

	return c
}
