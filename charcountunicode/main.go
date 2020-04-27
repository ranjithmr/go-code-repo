package main

import (
	"fmt"
	"unicode"
	"bufio"
	"log"
	"io"
	"os"
	"unicode/utf8"
)

func main(){
	countchar := make(map[rune]int)
//	countlen := make(map[rune]int)
//	countdig := make(map[rune]int)
	var utflen [utf8.UTFMax + 1]int
	invalid := 0
	input := bufio.NewReader(os.Stdin)
	for { 
	    r, n, err := input.ReadRune()
		if err == io.EOF {
		break
		}
	if err != nil { 
	log.Fatalln(err)
	}
//	if unicode.IsLetter(r) {
//		countlen[r]++
//	}
//	if unicode.IsDigit(r) {
//		countdig[r]++
//	}
	if r == unicode.ReplacementChar && n == 1 {
		invalid++
		continue
	}
	countchar[r]++
	utflen[n]++
	break
	}
//	fmt.Printf("length count %q\t", countlen)
//	fmt.Printf("length count %q\t", countdig)
	fmt.Printf("chararter count : %q\t", countchar)
	fmt.Printf("utf is %q\t", utflen)
	if invalid > 0 {
	fmt.Printf("length count %d\t", invalid)
	}
}
