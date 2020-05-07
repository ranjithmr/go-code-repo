package main

import (
	"html/template"
	"log"
	"net/http"
	//	"os"
)

var tpl *template.Template
var err error

func init() {
	tpl, err = template.ParseFiles("templates/about.gohtml", "templates/experience.gohtml", "templates/education.gohtml")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	http.HandleFunc("/", base)
	http.HandleFunc("/experience", experience)
	http.HandleFunc("/education", education)
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func base(w http.ResponseWriter, r *http.Request) {
	err = tpl.ExecuteTemplate(w, "about.gohtml", nil)
	if err != nil {
		log.Fatal(err)
	}
}
func experience(w http.ResponseWriter, r *http.Request) {
	err = tpl.ExecuteTemplate(w, "experience.gohtml", nil)
	if err != nil {
		log.Fatal(err)
	}
}
func education(w http.ResponseWriter, r *http.Request) {
	err = tpl.ExecuteTemplate(w, "education.gohtml", nil)
	if err != nil {
		log.Fatal(err)
	}
}
