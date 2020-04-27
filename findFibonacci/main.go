package main

import (
	"fmt"
	"custompackages/fibonacci"
)

func main() {
	fmt.Println(fibonacci.FibRecursive(7))
	fmt.Println(fibonacci.FibNormal(14))
}
