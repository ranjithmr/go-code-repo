package main

import (
	"fmt"
	"sort"
	"sync"
	"time"
)

func main() {
	var T int
	fmt.Scan(&T)
	var wg sync.WaitGroup

	for i := 0; i < T; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			var N int
			fmt.Scan(&N)
			teamG := make([]int, N)
			for i := 0; i < N; i++ {
				fmt.Scan(&teamG[i])
			}
			teamOP := make([]int, N)
			for i := 0; i < N; i++ {
				fmt.Scan(&teamOP[i])
			}
			start := time.Now()
			sort.Ints(teamG)
			sort.Ints(teamOP)
			count := 0
			for _, v := range teamG {
				for i, _ := range teamOP {
					fmt.Println(teamOP)
					if v <= teamOP[i] {
						break
					} else {
						count++
						teamOP = teamOP[(i + 1):]
						break
					}
				}
			}
			fmt.Println(count, time.Since(start))
		}()
		wg.Wait()
	}
}
