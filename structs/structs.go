package structs

import (
	"fmt"
)

type User struct {
	FirstName string // fields and fields types
	LastName  string
	Birthdate string
}

func (u *User) clearUserName() {
	u.FirstName = ""
	u.LastName = ""
}

func Structs() {

	// var appUser User

	appUser := User{
		FirstName: getUserData("Please enter the first name: "),
		LastName:  getUserData("Please enter the last name: "),
		Birthdate: getUserData("Please enter the birthdate (dd/mm/yyyy): "),
	}

	// outputUserDetails(appUser) ///
	// outputUserDetails(&appUser) /// address of appUser
	appUser.outputUserDetails() // calling the method
	appUser.clearUserName()
	appUser.outputUserDetails()
}

// passing structs as arguments
// func outputUserDetails(u User) {
// 	fmt.Println(u.FirstName, u.LastName, u.Birthdate)

// }
// func outputUserDetails(u *User) { ///dereferencing the value
// 	//longer approach to derefencing a pointer
// 	fmt.Println((*u).FirstName, u.LastName, u.Birthdate)

// }

// /methods
func (u User) outputUserDetails() {

	fmt.Println(u.FirstName, u.LastName, u.Birthdate)

}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
