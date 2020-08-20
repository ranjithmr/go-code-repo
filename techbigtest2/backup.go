package main

import (
	"fmt"
	"sort"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	var T int32
	fmt.Scan(&T)
	for i := 0; i < int(T); i++ {
		wg.Add(1)
		go func(i int) {
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
			count := 0
			for _, v := range teamG {
				for i, _ := range teamOP {
					if v > teamOP[i] {
						count += 1
						teamOP = append(teamOP[:i], teamOP[(i+1):]...)
						break
					}
				}
			}
			fmt.Println(count, time.Since(start))
			wg.Done()
		}(i)
		wg.Wait()
	}
}
