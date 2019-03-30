package main

import "fmt"

func main() {
	x := reverse([]int{1, 2, 3, 4, 5})
	fmt.Println(x)
	xx := reverseRec([]int{1, 2, 3, 4, 5, 6})
	fmt.Println(xx)
}

// Reverse numbers using iterative method
func reverse(nums []int) []int {
	for i, j := 0, len(nums)-1; i < j; i++ {
		nums[i], nums[j] = nums[j], nums[i]
		j--
	}
	return nums
}

// Reversomg numbers using recursive method
func reverseRec(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	newnums := []int{nums[len(nums)-1]}
	newnums = append(newnums, reverseRec(nums[:len(nums)-1])...)
	return newnums

}
