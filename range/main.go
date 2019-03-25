package main

import "fmt"

func main() {
	// composite literal of slice
	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	// Print all the value from slice using range
	for i, v := range xi {
		fmt.Println("Index ", i, "Value ", v)
	}

	fmt.Println("----------------------------")
	// composite literal of map
	mp := map[string]int{"Foo": 1, "Bar": 2}

	// Printing all the elements from map
	for key, value := range mp {
		fmt.Println("Key : ", key, "Value : ", value)
	}

	fmt.Println("------------ONLY KEYS----------------")
	// Fatching only keys
	for key := range mp {
		fmt.Println(key)
	}

	fmt.Println("----------ONLY VALUES------------------")
	// Fatching only values
	for _, value := range mp {
		fmt.Println(value)
	}

}
