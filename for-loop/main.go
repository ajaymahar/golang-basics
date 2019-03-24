package main

import "fmt"

// Program to reverse string
// So the Idea is that to convert string into bytes representations
// and then perform swaping operations
// then print complete byte with string conversion
func main() {
	str := "You Have a String"
	bs := []byte(str) //converting string into byte so we can assign values. we can't achive same with string bcz it's immutable

	// Here i = 0 and j = len(str) -1
	for i, j := 0, len(bs)-1; i < j; i++ {
		bs[i], bs[j] = bs[j], bs[i] //swapping first char with last char
		//j--
	}
	fmt.Println(string(bs))
	//fmt.Println(bs)   // byte representaion of string
}
