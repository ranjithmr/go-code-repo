package bank

import "fmt"

var deposits = make(chan int)
var balances = make(chan int)
var withdraw = make(chan int)

func Deposit(amount int) {
	deposits <- amount
}

func Balance() {
	fmt.Println(<-balances)
}

func Withdraw(amount int) {
	//	res := <-balances
	//if amount <= res {

	withdraw <- amount
	//	}
	//	fmt.Println("Insufficient Funds")
}
func init() {
	go teller()
}

func teller() {
	var balance int
	for {
		select {
		case amount := <-deposits:
			balance += amount
			fmt.Printf("successful. total balance: %d\n", balance)
		case amt := <-withdraw:
			if amt <= balance {
				balance -= amt
				fmt.Printf("Successful. total balance: %d\n", balance)
			} else {
				fmt.Println("Insufficient Funds")
			}
		case balances <- balance:
		}
	}
}
