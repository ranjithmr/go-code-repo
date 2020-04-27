package findtwothatsum


func FindSum(num []int ,number int)(int, int) {
	for i, n := range num {
	  for j, n1 := range num {
		if i == j {
		continue
		}
		if n + n1 == number {
		return n, n1
	 	}
	 }
	}
	return 0, 0
}
