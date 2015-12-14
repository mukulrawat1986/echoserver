// This version does the same work as server1, but it also counts
// the number of requests

package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex // mutex to lock the counting portion
var count int     // global counter

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echos the Path component of the requested url
func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

// counter echoes the number of requests made so far
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprint(w, "Count %d\n", count)
	mu.Unlock()
}
