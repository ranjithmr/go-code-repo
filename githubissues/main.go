package main

import (
	"log"
	"net/http"
	"os"
	"html/template"
	"githubissues/gthub"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	result, err := gthub.SearchIssues(os.Args[1:])
	if err != nil {
        log.Fatalln(err)
        }
	tpl := template.Must(template.ParseFiles("templates/gthbiss.gohtml"))
	err = tpl.Execute(w, result)
	if err != nil {
	log.Fatalln(err)
	}
	})
	log.Fatal(http.ListenAndServe(":8080", nil))

}
