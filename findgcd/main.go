package main

import (
	"fmt"
	"custompackages/gcd"
)

func main() {

	fmt.Println(gcd.DividerRec(16, 24))
	fmt.Println(gcd.DividerNor(15, 30))

}
