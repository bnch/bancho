package web

import (
	"fmt"
	"net/http"
	"os"
)

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
	avatarEngine = makeAvatarServer()

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
