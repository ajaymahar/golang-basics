package main

import "fmt"

const (
	//Ignoring initial iota value
	_  = iota
	kb = 1 << (iota * 10) // geting binary equalent of KB
	mb = 1 << (iota * 10) // geting binary equalent of MB
	gb = 1 << (iota * 10) // geting binary equalent of GB
	tb = 1 << (iota * 10) // geting binary equalent of TB
)

func main() {
	fmt.Println("Size in Byte \t Binary")
	fmt.Println("----------------------------")
	fmt.Printf("%d\t\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t\t%b\n", mb, mb)
	fmt.Printf("%d\t\t%b\n", gb, gb)
	fmt.Printf("%d\t%b", tb, tb)

}
