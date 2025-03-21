package functions

import "fmt"

func Functions() {
	fmt.Println("Functions")

	numbers := []int{1, 2, 3, 4}
	doubled := doubleNumbers(&numbers)

	fmt.Println(doubled)

}

func doubleNumbers(numbers *[]int) []int {
	dNumbers := []int{}
	for _, val := range *numbers {

		dNumbers = append(dNumbers, double(val))
	}
	return dNumbers
}

func double(num int) int {
	return num * 2
}
