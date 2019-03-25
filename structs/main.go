package main

import "fmt"

// Creating new struct called Person
// It will contains firstName, lastName, and age fields
type Person struct {
	firstName, lastName string
	age                 int
}

// Creating Programming struct
type Programming struct {
	isPython, isGo, isJava bool
}

// Creating new struct called
type Student struct {
	Person
	rollNumber uint
	Programming
}

func main() {
	// Defining composite literals
	/*per := Person{
		firstName: "Foo",
		lastName: "Bar",
		age:	42,
	}
	lang := Programming{
		isPython: true,
		isGo: true,
		isJava: false,
	}*/
	//fmt.Println(per.firstName, per.lastName, per.age)
	foobar := Student{
		Person: Person{
			firstName: "Foo",
			lastName:  "Bar",
			age:       42,
		},
		//Person: per,
		rollNumber: 1,
		Programming: Programming{
			isPython: true,
			isGo:     true,
			isJava:   false,
		},
		//Programming: lang,
	}
	kinley := Student{
		Person: Person{
			firstName: "Kinley",
			lastName:  "Minerals",
			age:       32,
		},
		rollNumber: 2,
		Programming: Programming{
			isPython: false,
			isGo:     false,
			isJava:   true,
		},
	}

	fmt.Println("-----------------------------------")
	fmt.Println("Name: ", foobar.firstName, foobar.lastName, "\nAge: ", foobar.age, "\nKnown Languages: \n Python\t  Go\tJava\n", foobar.isPython, "\t", foobar.isGo, "\t", foobar.isJava, "\n")
	fmt.Println("-----------------------------------")
	fmt.Println("Name: ", kinley.firstName, kinley.lastName, "\nAge: ", kinley.age, "\nKnown Languages: \n Python\t  Go\tJava\n", kinley.isPython, "\t", kinley.isGo, "\t", kinley.isJava)
}
