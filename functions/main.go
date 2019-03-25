package main

import "fmt"

func main() {
	//Basic function calling
	sayHello()
	// calling parameter function
	// Argument greet, name
	// Also called call by value
	greetings("Welcome, ", "Foo")

	// Calling parametized function
	// Argument greet, name
	// Call by refrence
	greet := "Welcome"
	name := "Bar"
	greeting(&greet, &name)
	fmt.Println("call by ref => ", name)

	//function with variadic parameters with int values
	sum(1, 2, 3, 4, 5)

	// function with variadic parameter and return statement
	total := total(1, 2, 3, 4, 5, 6)
	fmt.Println("Veriadic with return =>", total)

	//function with variadic param returnting local pointer
	ptotal := psum(1, 2, 3, 4)
	fmt.Println("Returning local pointer =>", *ptotal)

	// Name return value
	ntotal := nsum(5, 4, 3, 2, 1)
	fmt.Println("Name return function => ", ntotal)

	// calling function which returns multi values
	//result, err := divide(5.0, 3.0)
	result, err := divide(5.0, 0.0)
	//checking error
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Func return multi values =>", result)

}

//Basic function
func sayHello() {
	fmt.Println("Basic func => Hello Go")
}

//Function with parameters
//greet string, name string are same as greet, name string (with go syntatical suger)
func greetings(greet, name string) {
	fmt.Println("Call by value =>", greet, name)
}

// Call by refrence function
func greeting(greet, name *string) {
	fmt.Println(*greet, *name)
	*name = "Bizz"
}

// Variadic function
// accept unlimited number of int values
// variadic parameter must be the last one
// like sum(name string, nums ...int)
// invalid sum(nums ...int, name string)
func sum(nums ...int) {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("Simple Variadic =>", sum)
}

// Variadic function with return value
// This is same as above function except we are returning the value to the caller

func total(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

// Function which return local variable as pointer
// Buty of Golang not present in lot of languages
// when go sees we are passing address of varialbe it just promot the variable to the heap, so caller can use that
func psum(nums ...int) *int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	//sum = 90
	return &sum
}

// Name return value not very popular in golang comunity
func nsum(nums ...int) (result int) {
	for _, num := range nums {
		result += num
	}
	return
}

// Function return multi values

func divide(num1, num2 float64) (float64, error) {
	if num2 == 0.0 {
		return 0.0, fmt.Errorf("Can't divide by zero!")
	}
	return num1 / num2, nil
}
