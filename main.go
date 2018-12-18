package main

import (
	"github.com/goji/httpauth"
	"github.com/zenazn/goji"
	"net/http"
	"os"
)

func main() {
	u := os.Getenv("USER")
	p := os.Getenv("PASSWORD")
	if len(u) > 0 && len(p) > 0 {
		goji.Use(httpauth.SimpleBasicAuth(u, p))
	}
	// myHandler requires HTTP Basic Auth to access
	goji.Get("/*", http.FileServer(http.Dir("./docs")))
	goji.Serve()
}
