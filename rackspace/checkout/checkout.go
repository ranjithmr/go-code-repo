package checkout

import (
	"fmt"
	"rackspace/checkout/product"
)

type Cart struct {
	Basket    map[string]float64
	Catalogue []product.Item
}

func (c Cart) AddItem(item string) {
	c.Basket[item]++
}

func (c Cart) RemoveItem(item string) {
	_, ok := c.Basket[item]

	if ok {
		delete(c.Basket, item)
	} else {
		fmt.Println("Item not present in your basket")
	}
}

func Checkout(code ...string) {

	Items := product.Menu()
	sc := Cart{
		Basket:    make(map[string]float64),
		Catalogue: Items,
	}
	for _, it := range code {
		sc.AddItem(it)
	}
	var total float64
	fmt.Println("==========Basket Before Applying Offer===========")
	fmt.Println("\tProduct\tPrice")
	fmt.Println("==================================================")

	for code, quantity := range sc.Basket {
		for _, p := range sc.Catalogue {
			if code == p.Productcode {
				total += quantity * p.Price
				for quantity > 0 {
					fmt.Print("\t", code, "\t\t$", p.Price, "\n")
					quantity--
				}
			}
		}

	}
	fmt.Println("=================================================")
	fmt.Println("Before Discount: ", total)

	fmt.Println("=================================================")

	applydiscount := Promotion(sc, total)
	fmt.Println("=================================================")
	fmt.Println("After Discount: ", applydiscount)
	fmt.Println("=================================================")

}
