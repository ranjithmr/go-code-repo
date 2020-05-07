package main

import (
	"fmt"
	_ "sort"
	"time"
)

func main() {
	start := time.Now()
	var N int
	fmt.Scan(&N)

	ingredients := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&ingredients[i])
	}

	units := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&units[i])
	}
	for i, _ := range units {
		units[i] = units[i] / ingredients[i]
	}
	res := units[0]
	for _, n := range units[1:] {
		if n < res {
			res = n
		}
	}
	//sort.Ints(units)
	fmt.Println(res, time.Since(start))
}
