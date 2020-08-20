package main

import "go-code-repo/banker/bank"

func main() {

	bank.Deposit(200)
	bank.Withdraw(100)

	bank.Balance()
}
