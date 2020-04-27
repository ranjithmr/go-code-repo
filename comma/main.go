package main

import (
	"fmt"
//	"strings"
)

func main() {

	s := "12345678"
	b := comma(s)
	fmt.Println(b)
}

func comma(s string)string {
	n := len(s)
	if n <= 3 {
	return s
	}
	for i := 2; i < n; {
		s = s[:i] + "," + s[i:]
		i += 3
	}
	return s
}
