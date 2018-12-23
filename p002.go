package main

import "fmt"

//	Problem: By considering the terms in the Fibonacci sequence whose values do not exceed four million, find the sum of the even-valued terms.
func main() {
	res := 0
	nextFubonacciNumber := fubonacci()
	for curr := 0; curr < 4000000; curr = nextFubonacciNumber() {
		if curr%2 == 0 {
			res += curr
		}
	}
	fmt.Println(res)
}

func fubonacci() func() int {
	x, y := 0, 1
	return func() int {
		res := y + x
		x = y
		y = res
		return res
	}
}
