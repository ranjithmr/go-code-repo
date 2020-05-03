package main

import (
	"fmt"
	"github.com/google/uuid"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", mycookie)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatal(http.ListenAndServe(":8081", nil))

}

func mycookie(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	cookie, err := r.Cookie("seemycookie")
	if err != nil {
		id := uuid.New()
		cookie = &http.Cookie{
			Name:  "seemycookie",
			Value: id.String(),
			//Secure: true,
			HttpOnly: true,
			MaxAge:   60,
		}
		http.SetCookie(w, cookie)
	}
	fmt.Println(cookie)
}
