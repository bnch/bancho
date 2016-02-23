package web

import (
	"compress/gzip"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/bnch/bancho/avatarserver"
	"net/http"
	"strings"
)

var frontendEngine *gin.Engine

// ConnectionHandler is a very basic connection handler, used to handle the
// different hosts of the osu! server, and forward requests to the right engine.
type ConnectionHandler struct{}

func (c ConnectionHandler) serveHTTPReal(w http.ResponseWriter, r *http.Request) {
	switch r.Host {
	case "c.ppy.sh", "c1.ppy.sh":
		// Forward all requests to the bancho server to the bancho handler.
		BanchoConnectionHandler(w, r)
	case "a.ppy.sh":
		avatarserver.Serve(w, r)
	default:
		frontendEngine.ServeHTTP(w, r)
	}
}

func (c ConnectionHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
		c.serveHTTPReal(w, r)
		return
	}
	// Set up the chunked transfer.
	flusher, ok := w.(http.Flusher)
	if !ok {
		fmt.Println("expected http.ResponseWriter to be an http.Flusher")
		return
	}
	w.Header().Set("Content-Encoding", "gzip")
	gz := gzip.NewWriter(w)
	defer gz.Close()
	gzr := gzipResponseWriter{Writer: gz, ResponseWriter: w}
	c.serveHTTPReal(gzr, r)
	flusher.Flush()
}
