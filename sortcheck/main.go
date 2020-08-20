package main

import (
	"fmt"
	"sort"
	"time"
)

func main() {
	start := time.Now()
	ss := []int{5, 3, 4, 7, 8, 9, 8, 4, 28, 34, 48, 1, 3, 7}
	//	sort.Slice(ss, func(i, j int) bool { return ss[i] < ss[j] })
	sort.Ints(ss)
	fmt.Println(ss, time.Since(start))
}
