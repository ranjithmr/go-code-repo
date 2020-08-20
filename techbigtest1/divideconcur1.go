package main

import (
	"fmt"
	"sort"
	//	"sync"
	//	"runtime"
	"time"
)

func main() {
	//	var wg sync.WaitGroup
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
	//	wg.Add(1)
	//	go func() {
	//		defer wg.Done()
	for i := 0; i < len(units); i++ {
		if units[i] == 0 && ingredients[i] == 0 {
			continue
		} else if units[i] == 0 {
			fmt.Println(0)
			return
		} else if ingredients[i] == 0 {
			continue
		}

		units[i] = units[i] / ingredients[i]
	}
	//	}()
	//	wg.Wait()
	sort.Ints(units)
	for _, v := range units {
		if v > 0 {
			fmt.Println(v, time.Since(start))
			break
		}
	}
	//	fmt.Println(runtime.NumCPU())
}
