package main

import (
	"fmt"
	"os"
	"strconv"
	"log"
)

const (
	metertofeet = 3.28
	metertoinches = 39.37
)
func main() {
	var x float64
	num := os.Args[1:]
	if len(num) == 0 {
	fmt.Print("enter a digit :")
        fmt.Scan(&x)
        MeterToFeetAndInch(x)
        } else {

	for _, digit := range os.Args[1:]{
	v, err := strconv.ParseFloat(digit, 64)
	if err!= nil {
		log.Fatalln(err)
	}
	 MeterToFeetAndInch(v)
	}
	}
}

func MeterToFeetAndInch(m float64) {

	fmt.Printf("%f meters = %f Feet and %f Inches\n ", m, m * metertofeet, m * metertoinches)
	
}
