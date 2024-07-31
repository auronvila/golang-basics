package main

import "fmt"

func main() {
	age := 19
	ageP := &age

	fmt.Println("age before mutating the age val", *ageP)
	// not creating a copy of the age but passing the reference created with the pointer (ageP)
	editAgeToAdultYears(ageP)

	fmt.Println("age after mutating the val", *ageP)
	//fmt.Println("Adult age", adultYears)
}

func editAgeToAdultYears(age *int) {
	//return *age - 18
	*age = *age - 18
}
