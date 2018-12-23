package main

import (
	"project-euler-go/utils"
	"strconv"
)

//	Find the largest palindrome made from the product of two 3-digit numbers.
func main() {
	var allPalindrome []int
	for x := 100; x < 1000; x++ {
		for y := 100; y < 1000; y++ {
			holder := x * y
			if isPalindrome(holder) {
				allPalindrome = append(allPalindrome, holder)
			}
		}
	}
	println(utils.SliceBiggest(allPalindrome))
}

func isPalindrome(i int) bool {
	t := strconv.Itoa(i)
	length := len(t)
	for i := 0; i < length/2; i++ {
		if t[i] != t[length-i-1] {
			return false
		}
	}
	return true
}
