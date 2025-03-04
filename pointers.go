package main

import "fmt"

func pointers() {
	age := 25

	agePointer := &age        // get the address of age
	getAdultYears(agePointer) // pass the pointer to the function
	fmt.Println(age)
}

func getAdultYears(age *int) int {
	// return *age - 18

	*age = *age - 18 //dereference

	return *age
}
