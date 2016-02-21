package web

import (
	"github.com/gin-gonic/gin"
)

func makeFrontend() *gin.Engine {
	c := gin.Default()

	c.GET("/", func(c *gin.Context) {
		c.String(200, "I like memes")
	})
	c.GET("/web/bancho_connect.php", func(c *gin.Context) {
		c.String(200, "us")
	})

	return c
}
