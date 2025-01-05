package main

import (
	"fmt"
	"math/rand"
)

func main() {
	const cut float64 = 15
	var stakeBankBalance float64 = 0
	var initialInvestment float64
	fmt.Print("Deposit money in the bank: ")
	fmt.Scanln(&initialInvestment)

	fmt.Printf("You started with an investment of Rs: %f\n", initialInvestment)

	for {
		fmt.Printf("How much you wanna bet: ")
		var betAmount float64 = 0
		var choice, randomChoice int
		fmt.Scanln(&betAmount)
		if betAmount > initialInvestment || betAmount <= 0 {
			fmt.Println("Invalid Amount")
			continue
		}
		fmt.Printf("Head or Tail; Press 0 for head and 1 for tail: ")
		fmt.Scanln(&choice)
		randomChoice = rand.Intn(2)
		if randomChoice == choice {
			fmt.Println("Hurray! You won!")
			initialInvestment += betAmount * (100 - cut) * 0.01
			stakeBankBalance -= betAmount * (100 - cut) * 0.01
		} else if randomChoice == 1-choice {
			fmt.Println("Oops! You loose!")
			initialInvestment -= betAmount
			stakeBankBalance += betAmount
		} else {
			break
		}
		fmt.Printf("Your balance: %f\n", initialInvestment)
	}
	fmt.Printf("Your balance: %f\n", initialInvestment)
	fmt.Printf("Stake bank balance: %f\n", stakeBankBalance)

}
