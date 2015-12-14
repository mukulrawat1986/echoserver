// Minimal echo server that returns the path component of the URL
// used to access the server.

package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Each request calls handler
	// the server runs the handler for each incoming request in a
	// separate goroutine
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
