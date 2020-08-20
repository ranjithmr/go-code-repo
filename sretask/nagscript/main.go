package main

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	url := "https://httpbin.org/post"

	if len(os.Args) != 2 {
		fmt.Println("Not the correct usage please use --help or -h to know more")
		os.Exit(2)
	}

	if os.Args[1] == "--help" || os.Args[1] == "-h" {
		fmt.Println(` 
 This script is used to place the servers in maintenance in Nagios.
 Adjust date and time in json file and pass it as argument to script.
 
 Usage ./pausealerts.go filename 
 `)
		os.Exit(1)
	}
	// reading from the command line
	file := os.Args[1]

	// Read the json file
	data, err := ioutil.ReadFile(file)
	myError(err)

	//Create a new request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	myError(err)
	req.Header.Set("Content-Type", "application/json")

	//creata a context with a timout of 100 milliseconds
	ctx, cancel := context.WithTimeout(req.Context(), 500*time.Millisecond)
	defer cancel()

	// Bind the new context into the request.
	req = req.WithContext(ctx)

	//Make the web call and return any error
	resp, err := http.DefaultClient.Do(req)
	myError(err)

	// close the response body
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	myError(err)

	//write the status code to stdout
	fmt.Println(resp.StatusCode)
	fmt.Println(string(body))
}

func myError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
