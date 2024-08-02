package main

import "fmt"

type Product struct {
	Id    int
	Title string
	Price float64
}

func main() {
	// 1)
	hobbies := [3]string{"basketball", "lifting", "coding"}

	// 2)
	firstHobby := hobbies[0]
	secondThirdHobby := hobbies[1:3]
	fmt.Println(hobbies)
	fmt.Println(firstHobby)
	fmt.Println(secondThirdHobby)

	// 3)
	secondHobby := hobbies[1]
	//var mySlice = []string{firstHobby, secondHobby}
	mySlice := []string{firstHobby, secondHobby}
	fmt.Println(mySlice)

	// 4)
	thirdHobby := hobbies[2]
	mySlice[0] = secondHobby
	mySlice[1] = thirdHobby
	fmt.Println(mySlice)

	// 5)
	dynamicArr := []string{"learning Go", "software engineering"}
	fmt.Println(dynamicArr)

	// 6)
	dynamicArr[1] = "software developer"
	fmt.Println(dynamicArr)
	dynamicArr = append(dynamicArr, "software engineering")
	fmt.Println(dynamicArr)

	// 7)
	dynamicProducts := []Product{
		{Id: 0, Title: "firstProduct", Price: 1.1},
		{Id: 1, Title: "secondProduct", Price: 1.2},
	}
	fmt.Println(dynamicProducts)

	dynamicProducts = append(dynamicProducts, Product{
		Id:    2,
		Title: "thirdProduct",
		Price: 1.3,
	})
	fmt.Println(dynamicProducts)

	discountPrices := []int{1, 2, 3, 4}
	prices := []int{5, 6, 7, 8, 9}

	// use the ... dots to take all the values inside the slice and append to the priceTotal one by one.
	pricesTotal := append(discountPrices, prices...)
	fmt.Println(pricesTotal)
}

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.
