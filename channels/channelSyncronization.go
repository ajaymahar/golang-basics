package main

/*
We can use channels to synchronize execution across goroutines.
Here’s an example of using a blocking receive to wait for a goroutine to finish
*/
import (
	"fmt"
	"time"
)

//Defining func called worker
func worker(notify chan bool) {
	fmt.Println("Working....")
	time.Sleep(time.Second)
	fmt.Println("Work is done")

	//Notify another go routine work is done
	notify <- true
}
func main() {
	//Create chanel
	c := make(chan bool)

	//calling worker func
	//Start a worker goroutine, giving it the channel to notify on.
	go worker(c)

	//Waiting for worker to complete
	//Block until we receive a notification from the worker on the channel.
	<-c

	fmt.Println("Got notification now runnig main routine")

}
