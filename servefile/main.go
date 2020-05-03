package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", image)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func image(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("/home/ranjithmr_reddy/go_projects/src/go-code-repo/servefile/about.gohtml"))
	err := tpl.Execute(w, "about.gohtml")
	if err != nil {
		log.Fatal(err)
	}
	//	http.ServeFile(w, r, "IMG-20191216-WA0013.jpg")
}
