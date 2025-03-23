package functions

import "fmt"

// custom function type
// type anotFunc func(int, []string, map[string][]int) ([]int, string)

func Functions() {
	fmt.Println("Functions")

	numbers := []int{1, 2, 3, 4}
	doubled := transformNums(&numbers, double)

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

// function as funtion parameters

func transformNums(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}
	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}
