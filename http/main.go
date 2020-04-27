 package main

import (
	"fmt"
	"net/http"
	"log"
	"sync"
)

var count int
var mu sync.Mutex

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL is : %v", r.URL.Path)
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "count is : %d\n", count)
	mu.Unlock()
}
