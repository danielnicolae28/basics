package arraysmapsslice

import "fmt"

func List() {
	fmt.Println("Arrays, slices and maps")
	var producNames = [4]string{"Book"}
	prices := [4]float64{20.2, 43.5, 9.9, 10.2}

	producNames[2] = "carpet"
	fmt.Println(prices[1])
	fmt.Println(producNames)

	// slice
	///start at index 1 until index 3(index 3 being excluded)
	//type of slices
	/*
		featuredPrice := prices[:3] start at the begining until 3
		featuredPrice := prices[1:] start at the begining until end

		can't  start at a negativ index or go beyond array length
		slices can be crated based on other slices
	*/
	featuredPrice := prices[1:3]
	fmt.Println(featuredPrice)

	fmt.Println(len(prices), cap(prices))
	fmt.Println(len(featuredPrice), cap(featuredPrice))
}
