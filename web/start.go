package web

import (
	"fmt"
	"git.zxq.co/ripple/go-bancho/packethandler"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

// Start begins the webserver for go-bancho, and starts processing requests to the server.
func Start(addrHTTP, addrHTTPS string) {
	certificateExist := true
	for _, i := range []string{"cert.pem", "key.pem"} {
		if _, err := os.Stat(i); os.IsNotExist(err) {
			certificateExist = false
			fmt.Println("cert and key files were not found in the current directory. go-bancho will not listen on https.")
			break
		}
	}
	http.HandleFunc("/", makeGzipHandler(ConnectionHandler))
	if certificateExist {
		go func() {
			fmt.Println("Starting to listen on " + addrHTTPS + "...")
			http.ListenAndServeTLS(addrHTTPS, "cert.pem", "key.pem", nil)
		}()
	}
	fmt.Println("Starting to listen on " + addrHTTP + "...")
	http.ListenAndServe(addrHTTP, nil)
}

// ConnectionHandler takes inbound connections to the server and makes a sensed response.
func ConnectionHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html; charset=utf-8")

	// In case we're not doing a request from the osu! client, display the "frontend".
	if r.Method != "POST" || r.UserAgent() != "osu!" {
		w.Write([]byte(StandardPage))
		return
	}

	// Log that we got a request.
	fmt.Printf("==> REQUEST (token: \"%s\")\n", r.Header.Get("osu-token"))

	// Get data from request body
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error while attempting to understand request:", err)
		return
	}

	// We're not using .Add() because it capitalizes the string automatically. We'd rather not.
	w.Header()["cho-protocol"] = []string{strconv.Itoa(ProtocolVersion)}
	w.Header().Add("Vary", "Accept-Encoding")

	// Handle the packet
	var output []byte
	newToken, err := packethandler.Handle(data, &output, r.Header.Get("osu-token"))
	if err != nil {
		fmt.Println("Error in bancho:", err)
	}

	// Finish it up.
	if newToken != "" {
		w.Header()["cho-token"] = []string{newToken}
	}
	w.Write(output)
}
