package main

import "fmt"

/*
Expected would be:

[0 0 0]
[1 2 3]
Lenght of Slice is:  3
[1 2 3 4 5 6]
Lenght of Slice is:  6
[1 2 3 4 5 6 0 0 0]
Lenght of Slice is:  9
S[1:3]: =>  [2 3]
S[:6] => [1 2 3 4 5 6]
S[4:] => [5 6 0 0 0]
S[:] => [1 2 3 4 5 6 0 0 0]
New Slice: =>  [5 4 3 2 1]
After assignment x values: => [42 4 3 2 1]
New Slice Values are: =>  [5 4 3 2 1]
New Slice Before performing delete; => [3 4 5 6 34 332 322 2342]
New Slice After performing delete; => [3 4 34 332 322 2342]
[[] [] [] []]
[[0] [1 2] [2 3 4] [3 4 5 6]]

*/

func main() {
	//Declare and created slice called "s" type int using make func
	s := make([]int, 3)

	// Slice are similar to arrays
	fmt.Println(s)

	// Assigning values to slice
	s[0] = 1
	s[1] = 2
	s[2] = 3

	//Printing the Slice
	fmt.Println(s)
	//Printing lenght of slice
	fmt.Println("Lenght of Slice is: ", len(s))

	// append function is unique to slices
	s = append(s, 4, 5, 6)
	fmt.Println(s)
	//Printing lenght of slice
	fmt.Println("Lenght of Slice is: ", len(s))

	// creating new slice called "s1"
	s1 := make([]int, 3)
	s = append(s, s1...)
	fmt.Println(s)
	//Printing lenght of slice
	fmt.Println("Lenght of Slice is: ", len(s))

	// Slice syntax starting
	fmt.Println("S[1:3]: => ", s[1:3])
	fmt.Println("S[:6] =>", s[:6])
	fmt.Println("S[4:] =>", s[4:])
	fmt.Println("S[:] =>", s[:])
	// Note: We can't access the elements with negative values
	//fmt.Println("S[4:2] => ", s[4:-1]) //err: invalid slice index -1 (index must be non-negative)go
	//values of starting index must be grater than ending index
	//fmt.Println("S[4:2] =>", s[4:2])  //err: invalid slice index: 4 > 2go

	// Concise slice defination
	ns := []int{5, 4, 3, 2, 1}
	fmt.Println("New Slice: => ", ns)

	// Declaring x and Assigning ns
	// x := ns
	//fmt.Println("X values: =>", x)

	// assigning x[0] with 42
	// x[0] = 42
	// fmt.Println("After assignment x values: =>", x)
	// fmt.Println("New Slice Values are: => ", ns)
	// After assignment x values: => [42 4 3 2 1]
	// New Slice Values are: =>  [42 4 3 2 1]

	// Use copy built in function to avoid above proble
	x := make([]int, len(ns))
	copy(x, ns)
	x[0] = 42

	//After assignment x values: => [42 4 3 2 1]
	//New Slice Values are: =>  [5 4 3 2 1]

	fmt.Println("After assignment x values: =>", x)
	fmt.Println("New Slice Values are: => ", ns)

	// Delete from slice
	// Delete 5,6 from slice
	dslice := []int{3, 4, 5, 6, 34, 332, 322, 2342}
	fmt.Println("New Slice Before performing delete; =>", dslice)

	dslice = append(dslice[:2], dslice[4:]...)
	fmt.Println("New Slice After performing delete; =>", dslice)

	// 2D slice or slice of slice
	sls := make([][]int, 4)
	fmt.Println(sls)
	//[[] [] [] []]
	//[[0] [1 2] [2 3 4] [3 4 5 6]]
	for i := 0; i < 4; i++ {
		innerLength := i + 1
		sls[i] = make([]int, innerLength)
		for j := 0; j < innerLength; j++ {
			sls[i][j] = i + j
		}
	}
	fmt.Println(sls)
}
