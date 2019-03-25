package main

import "fmt"

// Created struct of greeter
type greeter struct {
	msg  string
	name string
}

// func (r Receiver) identifier(param)(return(s)) {code}
//associating greeter struct to a function it's called method
func (g greeter) greet() {
	fmt.Println(g.msg, g.name)
}

func main() {
	// Declared var and assigned struc with values
	ref := greeter{
		msg:  "Welcome, ",
		name: "Foo",
	}
	//calling method of greeter struct
	ref.greet()
}
