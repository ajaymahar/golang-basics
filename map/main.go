package main

import "fmt"

func main() {
	// Creating map[key]value, where key type is string and value type is int
	mp := make(map[string]int)
	fmt.Println(mp)
	fmt.Println(len(mp))

	// Adding values to mp
	mp["Foo"] = 1
	mp["Bar"] = 2
	fmt.Println(mp)
	fmt.Println(len(mp))

	// accessing the value from map
	// map will return value, ok as int , bool repectivily
	value, ok := mp["Fowero"]
	fmt.Println(value)
	fmt.Println(ok)

	// you can use the blank identifier (_) in place of the usual variable for the value.
	// If you just want to check element is present on map
	if _, ok := mp["sdf"]; !ok {
		fmt.Println("Key not found!")
	}
	// Common Go lang practice for checking the value is present or not
	if value, ok := mp["Foo"]; ok {
		fmt.Println("Key is present, here is the value : ", value)
	}

	// Deleting the element
	delete(mp, "Foo")
	fmt.Println(mp)

	// composite literal
	nmp := map[string]int{"Go": 1, "python": 2, "java": 3}
	fmt.Println(nmp)
	fmt.Println(len(nmp))

}
