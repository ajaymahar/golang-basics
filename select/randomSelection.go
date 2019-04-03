package main

import (
	"fmt"
	"time"
)

func dbServer1(ch chan string) {
	ch <- "From DB1"
}

func dbServer2(ch chan string) {
	ch <- "From DB2"
}

func main() {
	//Create Channel
	c1, c2 := make(chan string), make(chan string)
	go dbServer1(c1)
	go dbServer2(c2)

	time.Sleep(1 * time.Second)
	select {
	case res := <-c1:
		fmt.Println(res)
	case res := <-c2:
		fmt.Println(res)
	}
}
