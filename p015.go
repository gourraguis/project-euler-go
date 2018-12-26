package main

import (
	"fmt"
	"math/big"
)

//	Problem: How many such routes are there through a 20×20 grid (Lattice paths) ?
func main() {
	fmt.Println(latticePaths(20))
}

func latticePaths(size int) *big.Int {
	//	Shortest path diagrams, the number of paths is the central binomial coefficients
	//	Using big packages because the numbers can't be stored in an int64
	dividend, divisor := big.NewInt(1), big.NewInt(1)

	for i := 1; i <= size; i++ {
		divisor.Mul(divisor, big.NewInt(int64(i)))
	}
	for i := int64(1); i <= int64(2*size); i++ {
		dividend.Mul(dividend, big.NewInt(i))
	}
	return dividend.Div(dividend, divisor).Div(dividend, divisor)
}
