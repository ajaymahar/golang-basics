package main

import (
	"fmt"
	"sync"
)

//Create func that will print seriese of numbers

func printer(wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	wg.Done() // task is done
}

func main() {

	// Create wait group vars
	var wg sync.WaitGroup
	wg.Add(1) //adding count to wg
	go printer(&wg)

	wg.Wait() // waiting for task to complete
	fmt.Println("All Tasks are done!")
}
