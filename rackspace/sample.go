package main

import (
	"fmt"
	"rackspace/checkout"
)

func main() {
	fmt.Println("TEST 1")
	checkout.Checkout("CF1")
	fmt.Println()
	fmt.Println("TEST 2")
	fmt.Println()
	checkout.Checkout("CH1", "AP1", "AP1", "AP1")
	fmt.Println()
	fmt.Println("TEST 3")
	fmt.Println()
	checkout.Checkout("MK1", "AP1")

}
