package main

import "github.com/ajaykmahar/golang-programming/golang-programming/oop/employee"

func main() {
	//Create new employee
	//initialize with values
	e := employee.New("Ajay", "Mahar", 45, 20)
	//calling leaveremaining func
	e.LeaveRemaining()
}
