 package main

import (
	"fmt"
	"net/http"
	"log"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for i, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", i, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "Remote Add = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for i, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", i, v )
	}
}
