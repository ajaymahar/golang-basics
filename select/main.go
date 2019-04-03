package main

import (
	"fmt"
	"time"
)

//Create somework function
func somework(ch chan string) {
	time.Sleep(5 * time.Second)
	ch <- "Go routine work is done"
}
func main() {
	fmt.Println("Starting main func...")

	//Create chan
	ch := make(chan string)
	go somework(ch)

	for {
		time.Sleep(1 * time.Second)
		select {
		case msg := <-ch:
			fmt.Println(msg)
			return
		default:
			fmt.Println("data not recieved yet!")
		}
	}

}
