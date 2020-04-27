package fibonacci

func FibRecursive(num int)int {
	if num <= 1{
	return num
	}
	return FibRecursive(num-1) + FibRecursive(num-2)
}

func FibNormal(num int)int {
	if num <= 1{
        return num
        }

	nMinus2 := 0
	nMinus1 := 1
	var curN int
	for i := 2; i <= num; i++ {
		curN = nMinus1 + nMinus2
		nMinus2 = nMinus1
		nMinus1 = curN
	}
	return curN
}
