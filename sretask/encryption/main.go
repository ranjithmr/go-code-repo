package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
	"os"
)

func main() {
	var user string
	var db = make(map[string]string)
	password1 := "some user-provided password"

	// Generate "hash" to store from user password
	Pstore(user, password1, db)
	// After a while, the user wants to log in and you need to check the password he entered

	userPassword2 := "some user-provided password"

	// Comparing the password with the hash
	err := Pvalidate(user, userPassword2, db)
	if err != nil {
		log.Print("Invalid credentials")
		os.Exit(1)
	}
	fmt.Println("Authenication successful")
}
func Pvalidate(user, pass string, db map[string]string) error {
	err := bcrypt.CompareHashAndPassword([]byte(db[user]), []byte(pass))
	if err != nil {
		return err
	}
	return nil

}

func Pstore(user, pass string, db map[string]string) {

	// Generate "hash" to store from user password
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}

	password := string(hash)

	// Store this "hash" in a map
	//db := make(map[string]string)
	db[user] = password
}
