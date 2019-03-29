package main

import (
	"fmt"
	"math"
)

// Declaring rect type
type rect struct {
	width, height float64
}

//Declaring circle type
type circle struct {
	redius float64
}

//Declaring geometry interface
type geometry interface {
	area() float64
	perim() float64
}

//implementing methods for interface on rect type

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

//implementing methods for interface on circle type
func (c circle) area() float64 {
	return math.Pi * c.redius * c.redius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.redius
}

// Implementing measure func which accepts interface param
func measure(g geometry) {
	fmt.Println(g.area())
	fmt.Println(g.perim())
}
func main() {
	//Struct literals for rect and circle struc type
	r := rect{width: 10, height: 5}
	c := circle{redius: 7}

	//calling measure func
	measure(r)
	measure(c)
}
