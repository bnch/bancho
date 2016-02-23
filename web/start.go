package web

import (
	"fmt"
	"github.com/bnch/bancho/models"
	"github.com/bnch/bancho/packethandler"
	"github.com/bnch/bancho/avatarserver"
	"github.com/jinzhu/gorm"
	"net/http"
	"os"
)

var db gorm.DB

// Start begins the webserver for bancho, and starts processing requests to the server.
func Start(addrHTTP, addrHTTPS string) {

	certificateExist := true
	for _, i := range []string{"cert.pem", "key.pem"} {
		if _, err := os.Stat(i); os.IsNotExist(err) {
			certificateExist = false
			fmt.Println("cert and key files were not found in the current directory. bancho will not listen on https.")
			break
		}
	}

	setUpTemplates()
	frontendEngine = makeFrontend()
	
	var err error
	err = avatarserver.SetUp()
	if err != nil {
		panic(err)
	}

	db, err = models.CreateDB()
	if err != nil {
		panic(err)
	}

	packethandler.SetUp()

	handler := &ConnectionHandler{}
	if certificateExist {
		go func() {
			fmt.Println("Starting to listen on " + addrHTTPS + "...")
			http.ListenAndServeTLS(addrHTTPS, "cert.pem", "key.pem", handler)
		}()
	}
	fmt.Println("Starting to listen on " + addrHTTP + "...")
	http.ListenAndServe(addrHTTP, handler)
}
