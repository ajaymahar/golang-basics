package main

import "fmt"

//Creating function call greet
func greet(msg string) {
	for i := 0; i < 20; i++ {
		fmt.Println(msg, i)
	}
}

func main() {
	//calling function normally
	greet("Hey, there!")

	//Calling function async
	go greet("First Go Routine")

	// Calling same function with another go routrine
	go greet("Second Go Routine")

	//Calling anonymous function async way
	go func(msg string) {
		fmt.Println(msg)
	}("Anonymous func")

	fmt.Scanln()
	fmt.Println("Done!")
}
