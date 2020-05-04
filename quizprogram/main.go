package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	files := os.Args[1]

	if len(os.Args) == 0 {
		fmt.Println("Incorrect Usage: go run main.go filename is the correct usage ")
	}

	file, err := os.Open(files)
	if err != nil {
		log.Fatalf("error opening the file: %s", err)
	}
	defer file.Close()
	r := csv.NewReader(file)
	line, err := r.ReadAll()

	result := modifier(line)
	score := quiz(result)
	fmt.Printf("You scored %d out of %d in this quiz\n", score, len(result))

}

type problem struct {
	question string
	answer   string
}

func modifier(ss [][]string) []problem {
	pp := make([]problem, len(ss))
	for i, v := range ss {
		pp[i] = problem{
			question: v[0],
			answer:   strings.TrimSpace(v[1]),
		}
	}
	return pp
}

func quiz(result []problem) int {
	ansChan := make(chan string)
	var ans string
	count := 0
	for i, p := range result {
		timer := time.NewTimer(5 * time.Second)
		fmt.Printf("Question #%d : %s = ", i+1, p.question)
		go func() {
			fmt.Scanf("%s\n", &ans)
			ansChan <- ans
		}()

		select {
		case <-timer.C:
			fmt.Println("time out")
			continue
		case ans = <-ansChan:
			if p.answer == ans {
				if i == len(result)-1 {
					fmt.Println("End")
					count++
				} else {
					fmt.Println("Good job. Next Question.")
					count++
				}
			} else {
				if i == len(result)-1 {
					fmt.Println("End")
				} else {
					fmt.Println("Common. You can do better.")
				}
			}
		}
	}
	return count
}
