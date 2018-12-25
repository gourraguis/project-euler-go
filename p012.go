package main

import (
	"fmt"
	"project-euler-go/utils"
)

//	Problem: What is the value of the first triangle number to have over five hundred divisors?
func main() {
	fmt.Println(HighlyDivisibleTriangularNumber(500))
}

func HighlyDivisibleTriangularNumber(divisors int) int {
	res := 1

	for i := 2; utils.DivisorsNumber(res) < divisors; i++ {
		res += i
	}

	return res
}
