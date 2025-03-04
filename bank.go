package main

import (
	"fmt"

	"github.com/danielnicolae28/basics/fileops"
)

func Bank() {

	for {

		fmt.Println("Bank Go")
		userInput := inputChoice()
		balance := 1000

		switch userInput {
		case 1:
			CheckBalance(balance)
		case 2:
			WithdrawMoney(balance)
		case 3:
			DepositMoney(balance)
		case 4:
			fmt.Println("Thank you for working with us!")
			return
		}

	}

}

func inputChoice() (userInputChoice int) {

	fmt.Println("1. Check Balance")
	fmt.Println("2. Withdraw money")
	fmt.Println("3. Deposit money")
	fmt.Println("4. Exit")

	fmt.Print("Your Choice : ")
	fmt.Scan(&userInputChoice)

	if userInputChoice > 4 {
		fmt.Print("Not a valid input,Please select again one of the above : ")
		fmt.Scan(&userInputChoice)
	}

	return

}

func CheckBalance(balance int) {

	fmt.Printf("Balance : %v \n", balance)

}

func WithdrawMoney(balance int) {

	var withdrawAmount int

	fmt.Println("Please enter the amount to withdraw : ")
	fmt.Scan(&withdrawAmount)

	if balance < withdrawAmount {
		fmt.Printf("You don't have enough  money, please enter a sum equal or lower than your balance : %v : ", balance)

	} else {
		balance -= withdrawAmount
		fileops.WriteBalanceToFile(balance)
		balanceFromFile, _ := fileops.ReadBalanceFromFile()
		fmt.Printf("Balance : %v \n", balanceFromFile)
	}
}

func DepositMoney(balance int) {

	var depositAmount int

	fmt.Println("Please enter the amount to deposit : ")
	fmt.Scan(&depositAmount)

	if depositAmount < 0 {
		fmt.Println("Please enter a positiv amount ! ")
	} else {
		balance += depositAmount
		fileops.WriteBalanceToFile(balance)
		balanceFromFile, _ := fileops.ReadBalanceFromFile()
		fmt.Printf("Your balance is : %v \n", balanceFromFile)
	}

}
