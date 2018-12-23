package main

import (
	"fmt"
	"math"
	"math/big"
)

//	Problem: What is the largest prime factor of the number 600851475143 ?
func main() {
	fmt.Println(largestPrimeFactor(600851475143))
}

func largestPrimeFactor(N int64) int64 {
	maxPrime := math.Sqrt(float64(N))
	for i := int64(maxPrime); i > 1; i-- {
		if big.NewInt(i).ProbablyPrime(20) && N%i == 0 {
			return i
		}
	}
	return N
}
