package main

import (
	"fmt"
	"bytes"
//	"os"
)

func main() {
	s := "12345678"
	res := comma(s)
	fmt.Println(res)
}

func comma(s string) string {
 	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range s {
	if i > 3 {
	buf.WriteString(",")
	}
	fmt.Sprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
} 
