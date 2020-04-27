package main

import (
	"fmt"
)

func main() {
	s := []string{"abcd", "3456", "ijkl", "mn", "mn", "abcd"}
	fmt.Println(removedup(s))
}

func removedup(s []string)[]string {
        var result []string
        values := make(map[string]int)
        for _, v := range s {
                values[v] = 1
        }
        for key, _ := range values {
                result = append(result, key)
        }
        return result
}

