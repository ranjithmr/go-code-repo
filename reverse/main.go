package main

import (
	"fmt"
	"reverse/reverse"
	"reverse/rightrotate"
	"reverse/leftrotate"
)

func main() {
	s := []int{9, 8, 7, 6, 5}
	rev := reverse.Reverse(s)
	rr := rightrotate.RightRot(s, 1)
	lr := leftrotate.LeftRot(s, 2)
	fmt.Println(s, rev)
	fmt.Println(s, rr)
	fmt.Println(s, lr)
}
