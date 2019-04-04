package employee

import "fmt"

//Creating employee type struct
//Scope would be within employee package
type employee struct {
	firstName, lastName      string
	totalLeaves, leavesTaken uint
}

//Create new method to work as constuctor
//Scope would be outside of the employee package
//Initialize all the emp fileds and return emp struct
func New(firstName, lastName string, totalLeaves, leavesTaken uint) employee {
	emp := employee{firstName, lastName, totalLeaves, leavesTaken}
	return emp
}

//Creating method of employee struct
//Which will print the leaves of emp
func (e employee) LeaveRemaining() {
	fmt.Println(e.firstName, e.lastName, "Having", (e.totalLeaves - e.leavesTaken), "Leaves Remaining...")
}
