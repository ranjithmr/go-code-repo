package main

import (
	"fmt"
	"sort"
	_ "sync"
	"time"
)

var start time.Time

func main() {
	start = time.Now()
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
	s := materials{
		units,
		ingredients,
	}
	(&s).powerpuff()
}

type materials struct {
	units       []int
	ingredients []int
}

func (m *materials) powerpuff() {
	count := 0
	//	tr := make([]int, len(units))
	// compare compares the two slices and returns true if units is less than ingredients otherwise retuens false

	if m.compare() == false {

		for i := 0; i < len(m.units); i++ {
			m.units[i] = m.units[i] / m.ingredients[i]

		}

		sort.Ints[m.units]
		count = m.units[0]
	} else {
		count = 1
	}
	fmt.Println(count)
}

func (m *materials) compare() bool {
	for i, v := range m.units {
		if v < m.ingredients[i] {
			return true
		}
	}
	return false
}
