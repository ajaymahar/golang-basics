package main

import "fmt"

func main() {

	// Declaration and assignment of num variable
	num := 0
	fmt.Print("Enter positive of number: ") // Printing msg to console
	fmt.Scanf("%d", &num)                   // accepting value from user
	for i := 1; i <= int(num); i++ {        //looping thorough from 1 to provided numbers

		if i%2 == 0 { // checking current number is even/odd
			fmt.Println(i, "is an even number") //printing msg
		} else {
			fmt.Println(i, "is odd number")
		}
	}
}
