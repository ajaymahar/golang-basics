package main

import "fmt"

func main() {
	//Declare variable named arr sized/lenght 5
	var arr [5]int
	fmt.Println("Array", arr) //Printing an array [0,0,0,0,0]

	arr[4] = 42

	fmt.Println("Array: ", arr)
	fmt.Println("Array[4]: ", arr[4])
	fmt.Println("Length of an array: ", len(arr))

	//Declaration and assignment of an array sized/lenght 7
	newArray := [7]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println("New Array: ", newArray)
	fmt.Println("Lenght New Array: ", len(newArray))

	//2D array example
	var twoDArray [3][4]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			twoDArray[i][j] = i + j
		}
	}
	fmt.Println("2D array: ", twoDArray)
	fmt.Println("2D array Lenth:  ", len(twoDArray))

}
