package gcd

func DividerRec(a, b int)int {
	if b == 0 {
	return a
	}
	a, b = b, a%b
	return DividerRec(a, b)
}

func DividerNor(a, b int)int {
	for b != 0 {
	a, b = b, a%b
	}
	return a
}
