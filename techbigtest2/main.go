package main

import (
	"fmt"
	"sort"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var T int
	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		wg.Add(1)
		go func(a int) {
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
			fmt.Println(count)
			wg.Done()
		}(i)
		wg.Wait()
	}
}
