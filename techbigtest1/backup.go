package main

import (
	"fmt"
	"sort"
)

func main() {
	var N int
	fmt.Scanf("%d\n", &N)

	ingredients := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&ingredients[i])
	}

	units := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&units[i])
	}
	powerpuff(ingredients, units)

}

func powerpuff(ingredients, units []int) {
	count := 0
	if check(units, ingredients) == false {
		for i := 0; i < len(units); i++ {
			units[i] = units[i] / ingredients[i]
		}
		sort.Ints(units)
		count = units[0]
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
