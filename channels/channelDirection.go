package main

/*
When using channels as function parameters,
you can specify if a channel is meant to only send
or receive values. This specificity increases
the type-safety of the program.
*/

import "fmt"

//Create func called ping
//Which will only recieve ping msg, can't send
func ping(p chan<- string, msg string) {
	p <- msg
}

//Create func called pong
// <-chan only send channel
// chan<- only recieve channel
func pong(ping <-chan string, pong chan<- string) {
	//passing msg from ping chan to pong chan
	msg := <-ping
	pong <- msg
}
func main() {
	//Create channel
	cping, cpong := make(chan string), make(chan string)

	//calling ping func
	go ping(cping, "Sending request...")

	//calling pong func
	go pong(cping, cpong)

	fmt.Println(<-cpong)

}
