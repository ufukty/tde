package router

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"strings"

	"github.com/pkg/errors"
)

func dumpRequest(r *http.Request) {
	var dump, err = httputil.DumpRequest(r, false)
	if err != nil {
		log.Println(errors.Wrap(err, "dumping request"))
	}
	log.Printf("%q", strings.ReplaceAll(strings.ReplaceAll(string(dump), "\r\n", " || "), "\n", " | "))
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	dumpRequest(r)
	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
}

func LastMatch(w http.ResponseWriter, r *http.Request) {
	dumpRequest(r)
	if r.URL.Path == "/" {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	} else {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	}
}

func Pong(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "pong")
}
