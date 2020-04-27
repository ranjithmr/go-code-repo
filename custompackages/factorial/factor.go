package factorial


func Factor(num int) []int {
	var res []int
	primes := []int{2, 3, 5, 7, 9, 11, 13, 15, 17, 19} 

	if num == 1 {
		res = append(res, num)
                return res
        }

	for _, p := range primes {
	for num % p == 0 {
		num = num / p
		res = append(res, p)
	}
	}
	return res
}

