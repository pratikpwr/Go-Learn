package main

import "fmt"

func main() {

	var withdraw int64
	var balance float64
	fmt.Scan(&withdraw)
	fmt.Scan(&balance)

	remBal := withdrawMoney(withdraw, balance)

	fmt.Printf("%.2f", remBal)

}

func withdrawMoney(withdraw int64, balance float64) (remBalance float64) {

	if withdraw%5 == 0 && balance > float64(withdraw) {
		remBalance = balance - (float64(withdraw) + 0.50)
	} else {
		return balance
	}
	return
}
