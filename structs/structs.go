package main

import (
	"fmt"
	"github.com/structs/user"
)

func main() {
	var newUser *user.User
	newUser = user.New()
	if newUser.FirstName == "" || newUser.LastName == "" || newUser.BirthDate == "" {
		panic("Values cannot be empty")
	}

	admin := user.NewAdmin("test@gmail.com", "111")
	admin.User = newUser
	fmt.Println(admin)

	newUser.OutputUserData()
	newUser.ClearUserName()
	newUser.OutputUserData()
}
