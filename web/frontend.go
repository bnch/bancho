package web

import (
	"github.com/gin-gonic/gin"
)

func makeFrontend() *gin.Engine {
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
	c.POST("/signup", signupHandler)

	c.GET("/web/bancho_connect.php", func(c *gin.Context) {
		c.String(200, "us")
	})

	return c
}
