// Package avatarserver serves avatars from the folder data/avatars.
package avatarserver

import (
	"fmt"
	"github.com/bnch/hsfwsc"
	"net/http"
	"os"
	"strconv"
)

// Serve responds to an HTTP request with an avatar if present (responds with default avatar otherwise)
func Serve(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s a.ppy.sh%s\n", r.Method, r.URL.Path)
	if len(r.URL.Path) < 1 {
		defaultAvatar(w, r, 403)
		return
	}
	pic := r.URL.Path[1:]
	picID, err := strconv.Atoi(pic)
	if err != nil {
		defaultAvatar(w, r, 200)
		return
	}
	if _, err := os.Stat("data/avatars/" + strconv.Itoa(picID) + ".png"); os.IsNotExist(err) {
		defaultAvatar(w, r, 403)
		return
	}
	http.ServeFile(w, r, "data/avatars/"+strconv.Itoa(picID)+".png")
}

func defaultAvatar(w http.ResponseWriter, r *http.Request, statusCode int) {
	//w.WriteHeader(statusCode)
	hsfwsc.ServeFile(w, r, "data/avatars/default.png", statusCode)
}
