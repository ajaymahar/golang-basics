package main

import (
	"fmt"
	"sync"
)

//Create incrementer function with waitgroup and mutex
func incrementer(wg *sync.WaitGroup, mutex *sync.Mutex, idx *int) {
	mutex.Lock()   // locking the idx var for race condition
	*idx++         //increment idx value
	mutex.Unlock() //Unlocing the var
	wg.Done()      //Notifing other go routines task it done
}

func main() {
	var idx int
	var m sync.Mutex
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)                     //adding counter for every go routines
		go incrementer(&wg, &m, &idx) //calling incrementer func
	}
	wg.Wait()        //waiting for go routines to complete there there task
	fmt.Println(idx) // printing idx value
}
