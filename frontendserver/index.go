package frontendserver

import (
	"github.com/gin-gonic/gin"
)

func indexGET(c *gin.Context) {
	serveTemplate("index_public", nil, 200, c)
}
