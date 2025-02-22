package main

import (
	"fmt"
	"math"
)

func InvestCalculator() {

	var investAmount int
	var expectReturn float64
	var years int

	fmt.Println("Invest Calculator")

	fmt.Print("Enter invest amount : ")
	fmt.Scan(&investAmount)
	fmt.Print("Enter return procent: ")
	fmt.Scan(&expectReturn)
	fmt.Print("Enter period: ")
	fmt.Scan(&years)

	investmentAmountReturned := float64(investAmount) * math.Pow(1+expectReturn/100, float64(years))

	fmt.Println(investmentAmountReturned)

}
