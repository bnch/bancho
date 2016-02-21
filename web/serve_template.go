package web

import (
	"github.com/gin-gonic/gin"
)

func serveTemplate(templateName string, p interface{}, status int, c *gin.Context) {
	c.Data(status, "text/html", []byte{})
	templates[templateName].ExecuteTemplate(c.Writer, "tpl", p)
}
