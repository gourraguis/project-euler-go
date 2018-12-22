package main

import (
	"strconv"
)

//	Find the largest palindrome made from the product of two 3-digit numbers.
//	This is not the most optimal solution, but it gets the job done and this code won't be ran that many times.
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
	println(biggestInArr(allPalindrome))
}

func biggestInArr(arr []int) int {
	res := arr[0]
	for _, v := range arr {
		if v > res {
			res = v
		}
	}
	return res
}

func isPalindrome(i int) bool {
	t := strconv.Itoa(i)
	length := len(t)
	for i := 0; i < length / 2; i++ {
		if t[i] != t[length - i - 1] {
			return false
		}
	}
	return true
}