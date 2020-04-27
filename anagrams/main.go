package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	s1 := "SABUTRA"
	s2 := "BUATSARA"

	s3 := compare(s1)
	s4 := compare(s2)

	if s3 == s4{
	fmt.Println("Two strings are Anagrams")
	} else {
	fmt.Println("They are different")
	}
}

func compare(s string)string {
	spl := strings.Split(s, "")
	sort.Strings(spl)
	return strings.Join(spl, "")
}

