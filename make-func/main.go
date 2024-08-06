package main

import "fmt"

type myMapType map[string]string

func (m myMapType) output() {
	fmt.Println(m)
}

func main() {
	var mySlice []string
	mySlice = append(mySlice, "Auron")
	mySlice = append(mySlice, "Vila")

	for i, val := range mySlice {
		fmt.Println("slice - Index", i)
		fmt.Println("slice - Val", val)
	}

	myMap := make(myMapType, 3)
	myMap["Hello"] = "World"
	myMap["Hello1"] = "World1"
	myMap["Hello2"] = "World2"
	//myMap.output()

	for index, val := range myMap {
		fmt.Println("Index", index)
		fmt.Println("Val", val)
	}

}
