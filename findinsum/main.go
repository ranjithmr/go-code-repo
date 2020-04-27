package main

import (
	"fmt"
	"custompackages/findtwothatsum"
)


func main() {
	s := []int{ 2, 5, 7, 8, 9 }
	fmt.Println(findtwothatsum.FindSum(s, 10))

}
