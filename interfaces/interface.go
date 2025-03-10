package interfaces

import "fmt"

// interfaces are named collections of method signature

type calculator interface {
	add() int
	mult() int
	substr() int
	divide() int
	//interface{} ///nested interface
}

type userInput struct {
	a, b int
}

// methods
func (u userInput) add() int {
	return u.a + u.b
}
func (u userInput) mult() int {
	return u.a * u.b
}
func (u userInput) substr() int {
	return u.a - u.b
}
func (u userInput) divide() int {
	return u.a / u.b
}

func Measure(c calculator) {

	fmt.Println(c.add())
	fmt.Println(c.mult())
	fmt.Println(c.substr())
	fmt.Println(c.divide())
}

func InterfaceLacture() {
	fmt.Println("Interfaces")

	var a int
	var b int

	fmt.Println("Enter first value")
	fmt.Scan(&a)
	fmt.Println("Enter second value")
	fmt.Scan(&b)

	result := userInput{a: a, b: b}

	Measure(result)
	/// because userInput struct type implement the calculator interface we
	///can use instance of the struct as argument to Measure function

	//switch type
	printSomething(true)
	result2 := add2(1, 2)
	fmt.Println(result2)
}

/// swith type interface
/*
checking the type of input interface

*/

func printSomething(v interface{}) {

	switch v.(type) {
	case int:
		fmt.Println("int")
	case float64:
		fmt.Println("float64")
	case bool:
		fmt.Println("bool")
	default:
		fmt.Println("default value")
	}
}

//// generics

func add2[T int | float64 | string](a, b T) T {

	return a + b
}
