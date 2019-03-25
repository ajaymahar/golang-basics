package main

import "fmt"

func main() {
	// Creating anonymous struct
	// first {name string} is defining fields of struct
	// second {name: "XYZ"} is for initializer
	anonymous := struct{ name string }{name: "Foo Bar"}
	fmt.Println(anonymous)

	// Creating copy struct just copy the value not the refrence
	anotherAnonymous := anonymous
	//Update name
	anotherAnonymous.name = "Fizz Bizz"
	//Verify the both names
	fmt.Println(anotherAnonymous)
	fmt.Println(anonymous)

	//passign the refrence
	//anonPtr.name will point to Fly Run
	//anonymous.name will point to Fly Run
	anonPtr := &anonymous
	anonPtr.name = "Fly Run"
	fmt.Println(anonPtr.name)
	fmt.Println(anonymous.name)
}
