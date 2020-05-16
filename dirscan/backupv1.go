package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"
)

func main() {
	start := time.Now()
	dir := os.Args[1]
	if len(dir) == 0 {
		fmt.Println("usage :go run main.go directoryname")
	}
	var files []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {

		if err != nil {
			log.Fatal(err)
			return nil
		}
		// For Skipping the directory
		if info.IsDir() {
			return nil
		}

		// addding files to the slice we declared above
		files = append(files, path)
		return err
	})
	if err != nil {
		log.Fatal(err)
	}
	// range over the files and passing each file to createhash function to create the hash value for the given file
	for _, file := range files {
		createhash(file)
	}
	// passing slice of struct to findDuplicates function to find the files with same hash value
	findDuplicates(rr)
	fmt.Println(time.Since(start))
}

type Result struct {
	filename string
	hash     string
}

var rr []Result

func createhash(file string) {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}
	myhash := hex.EncodeToString(h.Sum(nil))
	myval := Result{
		filename: file,
		hash:     myhash,
	}
	rr = append(rr, myval)
}

func findDuplicates(rr []Result) {
	for i := 0; i < len(rr)-1; i++ {
		var equal []string
		for j := i + 1; j < len(rr); j++ {
			if rr[i].hash == rr[j].hash {
				equal = append(equal, rr[j].filename)
				rr = append(rr[:j], rr[(j+1):]...)
			}
		}
		if len(equal) > 0 {
			equal = append(equal, rr[i].filename)
			fmt.Printf("These files are equal %v\n", equal)
		}
	}
}
