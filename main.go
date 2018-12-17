package main

import (
	"github.com/goji/httpauth"
	"github.com/zenazn/goji"
	"net/http"
)

func main() {
	goji.Use(httpauth.SimpleBasicAuth("user", "password"))
	// myHandler requires HTTP Basic Auth to access
	goji.Get("/*", http.FileServer(http.Dir("./docs")))
	goji.Serve()
}
