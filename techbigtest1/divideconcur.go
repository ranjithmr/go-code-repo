package main

import (
	"fmt"
	"sort"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
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
	start := time.Now()
	for i := 0; i < len(units); i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			if units[i] == 0 && ingredients[i] == 0 || ingredients[i] == 0 {
				return
			} else if units[i] == 0 {
				fmt.Println(0)
				return
			}
			units[i] = units[i] / ingredients[i]
		}(i)
	}
	wg.Wait()
	sort.Ints(units)
	for _, v := range units {
		if v > 0 {
			fmt.Println(v, time.Since(start))
			break
		}
	}
}
