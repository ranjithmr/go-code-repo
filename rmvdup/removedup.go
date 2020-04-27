package removedup

func Removedup(s []string)[]string {
	var result []string
	values := make(map[string]bool)
	for _, v := range s {
		values[v] = true
	}
	for key, _ := range values {
		result = append(result, key)
	}
	return result
}
