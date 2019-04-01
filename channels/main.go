package main

import "fmt"

func main() {
	//Creating bidirectional channel
	ch := make(chan string)

	// Create async anonymous function which sends data into channels
	go func() {
		ch <- "Data"
		ch <- "Foo"
		ch <- "Bar"
		defer close(ch)
	}()

	// Recieve data from channels
	msg := <-ch

	//Printing msg
	fmt.Println(msg)

	//Printing another msg
	fmt.Println(<-ch)

	//Range over channel
	for value := range ch {
		fmt.Println("Inside Range")
		fmt.Println(value)
	}
}
