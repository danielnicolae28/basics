package main

import (
	"fmt"
	"math"
)

func InvestCalculator() {
	const inflationRate = 2.5

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
	futureRealValue := investmentAmountReturned / math.Pow(1+inflationRate/100, float64(years))
	fmt.Printf("Return amount : %.2f\n", investmentAmountReturned)
	fmt.Printf("Real value after inflation: %.2f\n", futureRealValue)

}
