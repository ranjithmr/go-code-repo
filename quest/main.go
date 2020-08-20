package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "RANJITHHH"
	res := duplicate(s)
	if res == true {
		fmt.Println("there are repeated character")
	} else {
		fmt.Println("there are no repeated character")
	}
}

func duplicate(s string) bool {
	ss := strings.Split(s, "")
	for i := 0; i < len(ss)-1; i++ {
		for j := i + 1; j < len(ss); j++ {
			if ss[i] == ss[j] {
				return true
			}
		}
	}
	return false
}
