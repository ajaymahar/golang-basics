package main

import "fmt"

func main() {
	// calling sum function with fixed size array
	fmt.Println(sum([5]float64{1, -2, 3, 4.0, 5.34}))
}

func sum(nums [5]float64) float64 {
	// Declaration and Assignment of variable sum
	sum := 0.0
	for _, num := range nums { // loop through all the elements of nums float64 array
		sum += num // Adding each element one by one
	}
	return sum // Returning sum of all the elements
}
