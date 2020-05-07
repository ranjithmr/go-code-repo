package main

import (
	"fmt"
	"sync"
)

func main() {
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
	powerpuff(units, ingredients)

}

func powerpuff(units, ingredients []int) {
	var wg sync.WaitGroup
	count := 0
	//	tr := make([]int, len(units))
	// compare compares the two slices and returns true if units is less than ingredients otherwise retuens false
	//ack := make(chan bool)

	for check(units, ingredients) == false {
		count++
		for i := 0; i < len(units); i++ {
			wg.Add(1)
			go func(a int) {
				units[a] = units[a] - ingredients[a]
				wg.Done()
			}(i)
		}
		//}
		wg.Wait()
	}
	fmt.Println(count)
}

func check(units, ingredients []int) bool {

	for i, v := range units {
		if v < ingredients[i] {
			return true
		}
	}
	return false
}
