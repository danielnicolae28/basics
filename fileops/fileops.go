package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func WriteBalanceToFile(balance int) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile("balance.json", []byte(balanceText), 0644)
}
func ReadBalanceFromFile() (float64, error) {

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
