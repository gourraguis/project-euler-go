package main

import "fmt"

//	Problem: Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.
func main() {
	fmt.Println(sumSquareDiff(100))
}

func sumSquareDiff(limit int) int {
	return squareOfSums(limit) - sumOfSquares(limit)
}

func sumOfSquares(limit int) int {
	res := 0
	for i := 1; i <= limit; i++ {
		res += i * i
	}
	return res
}

func squareOfSums(limit int) int {
	res := 0
	for i := 1; i <= limit; i++ {
		res += i
	}
	return res * res
}
