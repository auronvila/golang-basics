package user

import (
	"fmt"
	"time"
)

type User struct {
	FirstName string
	LastName  string
	BirthDate string
	CreatedAt time.Time
}

type Admin struct {
	email    string
	password string
	*User
}

// receiver argument
func (user *User) OutputUserData() {
	fmt.Println(user)
}

// if you want to change the value of the struct take it as a pointer otherwise
// it will create a copy of it, and you will have an unexpected behaviour

func (user *User) ClearUserName() {
	user.FirstName = ""
	user.LastName = ""
}

func NewAdmin(email, password string) *Admin {
	return &Admin{
		email:    email,
		password: password,
		User: &User{
			// alternatively you can get the user properties as params
			FirstName: "Admin",
			LastName:  "AAA",
			BirthDate: "04/05/2004",
			CreatedAt: time.Now(),
		},
	}
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}

func New() *User {
	return &User{
		FirstName: getUserData("Please enter your first name: "),
		LastName:  getUserData("Please enter your last name: "),
		BirthDate: getUserData("Please enter your birthdate (MM/DD/YYYY): "),
		CreatedAt: time.Now(),
	}
}
