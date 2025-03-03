package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
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
		fmt.Printf("You don't have enough  money, please enter a sum equal or lower to your balance : %v : ", balance)

	} else {
		balance -= withdrawAmount
		writeBalanceToFile(balance)
		balanceFromFile, _ := readBalanceFromFile()
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
		writeBalanceToFile(balance)
		balanceFromFile, _ := readBalanceFromFile()
		fmt.Printf("Your balance is : %v \n", balanceFromFile)
	}

}

func writeBalanceToFile(balance int) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile("balance.json", []byte(balanceText), 0644)
}
func readBalanceFromFile() (float64, error) {

	data, err := os.ReadFile("balance.txt")

	if err != nil {
		return 1000, errors.New("failed to parse stored data")
	}
	balanceText := string(data)

	balance, err := strconv.ParseFloat(balanceText, 64)

	if err != nil {
		return 1000, errors.New("failed to convert stored data")
	}

	return balance, nil
}
