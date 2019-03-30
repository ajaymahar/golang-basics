package main

import (
	"fmt"
	"strings"
)

func WordCount(words string) map[string]uint {

	wordCounts := make(map[string]uint)
	/*listOfWords := strings.Split(words, " ")
	for _, word := range listOfWords {
		if _, ok := wordCounts[word]; ok {
			wordCounts[word]++
		} else {
			wordCounts[word] = 1
		}
	}*/
	for _, word := range strings.Fields(words) {
		wordCounts[word]++
	}
	return wordCounts
}
func main() {
	words := "This is my string, this is awesome"
	fmt.Println(WordCount(words))
}
