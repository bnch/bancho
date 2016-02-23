package avatarserver

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer(t *testing.T) {
	err := SetUp()
	if err != nil {
		t.Fatal(err)
	}

	err = ioutil.WriteFile("data/avatars/1336.png", []byte("testing"), 0644)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("GET", "http://a.ppy.sh/1336", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	Serve(w, req)

	if w.Body.String() != "testing" {
		t.Fatalf("expected GET /1336 to return 'testing', got %s instead", w.Body.String())
	}

	req, err = http.NewRequest("GET", "http://a.ppy.sh/2", nil)
	if err != nil {
		t.Fatal(err)
	}
	w = httptest.NewRecorder()
	Serve(w, req)

	if w.Code != 403 {
		t.Fatalf("expecting GET to an invalid picture to return error code 403, got %d instead", w.Code)
	}

	req, err = http.NewRequest("GET", "http://a.ppy.sh/random", nil)
	if err != nil {
		t.Fatal(err)
	}
	w = httptest.NewRecorder()
	Serve(w, req)

	if w.Code != 200 {
		t.Fatalf("expecting GET to a path with a string rather than a number to return status code 200, got %d instead", w.Code)
	}

	req.URL.Path = ""
	req.URL.RawPath = ""
	w = httptest.NewRecorder()
	Serve(w, req)

	if w.Code != 403 {
		t.Fatalf("expecting GET to invalid path to return code 403, got %d instead", w.Code)
	}
}
