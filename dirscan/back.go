package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	_ "io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func main() {
	dir := os.Args[1]
	if len(dir) == 0 {
		fmt.Println("usage :go run main.go directoryname")
	}
	err := filepath.Walk(dir, dirscan)
	if err != nil {
		log.Fatalln(err)
	}
	findDuplicates(rr)
	//fmt.Println(rr)
}
func dirscan(path string, info os.FileInfo, err error) error {

	var files []string
	if err != nil {
		log.Fatal(err)
		return nil
	}

	if info.IsDir() {

		return nil
	}
	files = append(files, path)
	for _, file := range files {
		createhash(file)
	}
	return nil
}

type Result struct {
	filename string
	hash     string
}

var rr []Result

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
