package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"crypto/sha512"
	"log"
	"path/filepath"
)


func main() {
	dir := os.Args[1]
	if len(dir) == 0 {
		fmt.Println("usage :go run main.go directoryname")
	}
	err := filepath.Walk(dir, checkdup)
	if err != nil {
                log.Fatalln(err)
        }
}

var files = make(map[[sha512.Size]byte]string)

func checkdup(path string, info os.FileInfo, err error) error {

	if err != nil {
		log.Fatal(err)
		return nil
	}
	
	if info.IsDir() {
        	return nil
        }
	
	data, err := ioutil.ReadFile(path)
        if err != nil {
                log.Fatalln(err)
		return nil
        }
        
	myhash := sha512.Sum512(data)
	        
	if v, ok := files[myhash]; ok {
        fmt.Printf("%q is a duplicate of %q\n", path, v)
        } else {
	files[myhash] = path
	}
	return nil
}
