package frontendserver

import (
	"os"

	"github.com/bnch/bancho/bindata"
	"github.com/bnch/bancho/models"
	"github.com/elazarl/go-bindata-assetfs"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var frontendFolderExists bool

// Make creates a gin engine able to respond properly to requests.
func Make() *gin.Engine {
	if finfo, err := os.Stat("frontend"); !os.IsNotExist(err) && finfo.IsDir() {
		frontendFolderExists = true
	}

	setUpTemplates()

	var err error
	db, err = models.CreateDB()
	if err != nil {
		panic(err)
	}

	c := gin.Default()

	if frontendFolderExists {
		c.Static("/static", "frontend/static")
	} else {
		c.StaticFS("/static", &assetfs.AssetFS{
			Asset:     bindata.Asset,
			AssetDir:  bindata.AssetDir,
			AssetInfo: bindata.AssetInfo,
			Prefix:    "frontend/static",
		})
	}

	c.GET("/", indexGET)

	c.GET("/signup", signupGET)
	c.POST("/signup", signupPOST)

	c.GET("/web/bancho_connect.php", func(c *gin.Context) {
		c.String(200, "us")
	})

	return c
}
