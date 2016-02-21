package web

import (
	"fmt"
	"github.com/bnch/bancho/packethandler"
	"io/ioutil"
	"net/http"
	"strconv"
)

// BanchoConnectionHandler takes inbound connections to the bancho server (c.ppy.sh) and makes a sensed response.
func BanchoConnectionHandler(w http.ResponseWriter, r *http.Request) {
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
