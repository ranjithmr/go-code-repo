package main

import (
	"fmt"
	"time"
	"os"
	"io"
	"io/ioutil"
	"net/http"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, urls := range os.Args[1:] {
		go fetch(urls, ch)
	}

	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	
	fmt.Printf("%.2fs elasped\n", time.Since(start).Seconds())
}

func fetch(urls string, ch chan<- string) {
	start := time.Now()
	res, err := http.Get(urls)
		if err != nil {
		ch <- fmt.Sprintf("%s\n", err)
		return
		}
	b, err := io.Copy(ioutil.Discard, res.Body)
	res.Body.Close()
		if err != nil {
		ch <- fmt.Sprintf("%s\n", err)
		return
		}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, b, urls)
}
