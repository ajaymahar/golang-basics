package main

import "fmt"

func main() {
	// macking buffered channel type string
	//bch can only contains 2 elemets
	bch := make(chan string, 2)

	bch <- "First Msg"
	bch <- "Second Msg"

	//fatal error: all goroutines are asleep - deadlock!
	//bch <- "Uncomment me and run it"
	close(bch) //closing the channel If not range over will throw an error

	//range over bChannels
	for value := range bch {
		fmt.Println(value)
	}

}
