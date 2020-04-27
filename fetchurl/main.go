package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"io/ioutil"
	"strings"
)
func main() {
	urls := os.Args[1:]
	if len(urls) < 1 {
		fmt.Println("not enough argumemts : specify a url")
	}
	for _, ur := range urls {
		if strings.HasPrefix(ur, "https://") {
		res, err := http.Get(ur)
	 		if err != nil {
			log.Fatalln(err)
			os.Exit(1)
			}
		bs, err := ioutil.ReadAll(res.Body) 
				if err != nil {
				log.Fatalln(err)
				}
			res.Body.Close()
			code := res.Status
				
			fmt.Printf("%s\n%s\n", code, bs)
		} else {
			uri := "https://" + ur
			res, err := http.Get(uri)
				if err != nil {
				log.Fatalln(err)
				os.Exit(1)
				}
			bs, err := ioutil.ReadAll(res.Body) 
				if err != nil {
                        	log.Fatalln(err)
                        	}
                        res.Body.Close()

			code := res.Status
			fmt.Printf("%s\n%s\n", code, bs)
	}
	}
}
