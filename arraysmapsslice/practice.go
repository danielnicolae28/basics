package arraysmapsslice

import "fmt"

func Practice() {
	fmt.Println("Practice")

	//1. create a new array that contains three hobbies you have. Output that array

	hobbies := [3]string{"gym", "coding", "reading"}

	fmt.Println(hobbies)
	//--------------------
	//2. Output more data about that array:
	//the first element(standalone)
	//the second and third combined as a new list

	fmt.Println(hobbies[0])
	homeHobbies := hobbies[1:3]
	fmt.Println(homeHobbies)
	//-------

	//3. Create a slice based on the first element that contains the first and second elements
	//  create that slice in two different ways

	hobbiesTwo := hobbies[:2]
	fmt.Println("hobbies two", hobbiesTwo)
	hobbiesThree := hobbies[0:2]
	fmt.Println(hobbiesThree)

	//------
	/// 4. Re-slice the slice from 3 and change it to contain the second and last element of the original array

	hobbiesTwo = hobbiesTwo[1:3] // we expanded the slice
	fmt.Println(hobbiesTwo)
	fmt.Println(len(hobbiesTwo), cap(hobbiesTwo))

	//--------
	///// 5. Create a dynamic array that contains your course goals
	goals := []string{"learn go", "practice"}

	fmt.Println(goals)
	//// 6 set the second goal to a different one and the add a third goalt o that existing array
	goals[1] = "practice gooooooo"
	newGoals := append(goals, "create project")
	fmt.Println(newGoals)

	///// 7 create a product struct with title id, price and create a dynamic list of products at least 2 and add a third

	type produc struct {
		title string
		id    int
		price float64
	}

	productList := []produc{
		{title: "hat", id: 534, price: 6.87}, {title: "jeans", id: 5332, price: 61.87},
	}
	productListUpdated := append(productList, produc{title: "shirt", id: 233, price: 7.934})
	fmt.Println(productListUpdated)

}
