package main

import "fmt"

var timeZone = map[string]int{
	"UTC": 0 * 60 * 60,
	"EST": -5 * 60 * 60,
	"CST": -6 * 60 * 60,
	"MST": -7 * 60 * 60,
	"PST": -8 * 60 * 60,
}

func offset(tz string) int {
	if seconds, ok := timeZone[tz]; ok {
		return seconds
	}
	fmt.Println("Unknown timezone ", tz)
	return 0
}

func main() {
	fmt.Println(offset("CST"))
	fmt.Println(offset("IST"))
}
