package main

import (
	"fmt"
	"math"
)

func InvestCalculator() {
	const inflationRate = 2.5

	fmt.Println("Invest Calculator")

	investAmount, expectReturn, years := inputDataInvest()
	investmentAmountReturned, futureRealValue := InvestReturnCalculations(investAmount, expectReturn, inflationRate, years)

	fmt.Printf("Return amount : %.2f\n", investmentAmountReturned)
	fmt.Printf("Real value after inflation: %.2f\n", futureRealValue)

}

func InvestReturnCalculations(investAmount int, expectReturn float64, inflationRate float64, years int) (investmentAmountReturned float64, futureRealValue float64) {
	investmentAmountReturned = float64(investAmount) * math.Pow(1+expectReturn/100, float64(years))
	futureRealValue = investmentAmountReturned / math.Pow(1+inflationRate/100, float64(years))
	return
}

func inputDataInvest() (int, float64, int) {

	var investAmount int
	var expectReturn float64
	var years int

	fmt.Print("Enter invest amount : ")
	fmt.Scan(&investAmount)
	fmt.Print("Enter return procent: ")
	fmt.Scan(&expectReturn)
	fmt.Print("Enter period: ")
	fmt.Scan(&years)

	if investAmount > 0 || years > 0 {
		investAmount = 1000
		years = 5
	}

	return investAmount, expectReturn, years

}
