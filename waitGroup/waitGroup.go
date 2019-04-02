package main

import (
	"fmt"
	"sync"
)

//Create Wait group
var wg sync.WaitGroup

//Create func called foo

func foo() {
	for i := 1; i < 10; i++ {
		fmt.Println("Foo", i)
	}
	wg.Done()
}

//Create func called bar

func bar() {
	for i := 1; i < 10; i++ {
		fmt.Println("Bar", i)
	}
	wg.Done()
}

func main() {
	//Adding wait group counter
	wg.Add(1)
	go foo()

	//Adding wait group counter
	wg.Add(1)
	go bar()

	//watiging for all the goroutines to finish there tasks
	wg.Wait()
}
