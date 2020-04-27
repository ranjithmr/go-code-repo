package main

import (
	"fmt"
	"os"
	"log"
	"bufio"
)


func main() {
	count := make(map[string]int)
	if len(os.Args) == 0 {
		countLines(os.Stdin, count)
	} else {
	for _, fi := range os.Args[1:] {
		f, err := os.Open(fi)
		if err != nil {
			log.Fatalln(err)
			continue
			}
		countLines(f, count)
		f.Close()
		}
	}
	for line, n := range count {
         if n > 1{
                fmt.Printf("Duplicate found - %s\t %d\n ", line, n)
		
		}
	}
}
func countLines(f *os.File, count map[string]int) {
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		count[scanner.Text()]++
	}
	for line, n := range count {
         if n > 1{
                fmt.Printf("Duplicate found in %v - %s\t %d\n ",f.Name(), line, n)

                }
        }

  }
