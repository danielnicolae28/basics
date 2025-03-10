package structs

import (
	"errors"
	"fmt"
)

type customString string //  custom type

type User struct {
	FirstName string // fields and fields types
	LastName  string
	Birthdate string
}

type Admin struct {
	email    string
	password string
	User     // structs embeding
}

func (u *User) clearUserName() {
	u.FirstName = ""
	u.LastName = ""
}

func NewAdmin(email, password string) *Admin {

	return &Admin{
		email:    email,
		password: password,
		User: User{
			FirstName: "Admin",
			LastName:  "admin",
			Birthdate: "---",
		},
	}

}

// constructor function
func NewUser(firstName, lastName, birthday string) (*User, error) {

	if firstName == "" || lastName == "" || birthday == "" {
		return nil, errors.New("empty input") // check the input
	}

	return &User{
		FirstName: firstName, // capitalize to be used in other package
		LastName:  lastName,
		Birthdate: birthday,
	}, nil
}

func Structs() {

	var appUser *User
	var appAdmin *Admin

	firstName := getUserData("Please enter the first name: ")
	lastName := getUserData("Please enter the last name: ")
	birthday := getUserData("Please enter the birthdate (dd/mm/yyyy): ")
	email := getUserData("Please enter the email: ")
	password := getUserData("Please enter the pass ")

	appUser, err := NewUser(firstName, lastName, birthday)
	appAdmin = NewAdmin(email, password)

	if err != nil {
		fmt.Println("error")
		return

	}

	// outputUserDetails(appUser) ///
	// outputUserDetails(&appUser) /// address of appUser
	appUser.outputUserDetails() // calling the method
	appUser.clearUserName()
	appUser.outputUserDetails()
	appAdmin.outputAdmin()
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

func (a Admin) outputAdmin() {
	fmt.Println(a.email, a.User.LastName, a.password)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
