package web

import (
	"github.com/gin-gonic/gin"
)

func makeAvatarServer() *gin.Engine {
	c := gin.Default()

	c.GET("/", func(c *gin.Context) {
		c.String(404, "")
	})

	return c
}
