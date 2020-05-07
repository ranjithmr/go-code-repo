package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	mymux := http.NewServeMux()
	mymux.HandleFunc("/", mainpage)
	mymux.HandleFunc("/testing", testpage)
	log.Fatal(http.ListenAndServe(":8080", mymux))

}

func mainpage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the main page")
}

func testpage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("this should not print"))
}
