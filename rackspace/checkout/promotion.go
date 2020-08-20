package checkout

import "fmt"

func Promotion(sc Cart, total float64) float64 {
	fmt.Println("==========Basket Before Appllying Offer===========")
	fmt.Println("\tProduct\t\tPrice")
	fmt.Println("==================================================")

	for code, quantity := range sc.Basket {
		for _, p := range sc.Catalogue {
			if code == p.Productcode {
				q := quantity
				for q > 0 {
					fmt.Print("\t", code, "\t\t$", p.Price, "\n")
					q--
				}
				if code == "CH1" {
					//	sc.Basket["MK1"]++
					fmt.Print("\tMK1\t\t$4.75\n")
					fmt.Print("\t\tCHMK\t-$4.75\n")
				}
				if code == "AP1" && quantity >= 3 {
					total -= quantity * 1.5
					for quantity > 0 {
						fmt.Print("\t\tAPPL\t-$1.5\n")
						quantity--
					}
				}
				if code == "CF1" {
					for quantity > 0 {
						sc.Basket["CF1"]++
						fmt.Print("\tCF1\t\t$11.23\n")
						fmt.Print("\t\tBOGO\t-$11.23\n")
						quantity--
					}
				}
				if code == "OM1" {
					ap, ok := sc.Basket["AP1"]
					if ok {
						for quantity > 0 && ap > 0 {
							total = total - 3
							quantity--
							ap--
							fmt.Print("\t\tAPOM\t-$3\n")
						}
					}
				}
			}
		}
	}
	_, ok := sc.Basket["CH1"]
	if ok {
		sc.Basket["MK1"]++
	}
	return total
}
