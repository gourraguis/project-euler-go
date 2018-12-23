package main

import (
	"fmt"
	"math/big"
)

//	Problem: What is the 10 001st prime number?
func main() {
	fmt.Println(nthPrime(10001))
}

func nthPrime(n int) int {
	currentPrime := 0
	for i := 0; ; i++ {
		if big.NewInt(int64(i)).ProbablyPrime(20) {
			currentPrime++
			if currentPrime == n {
				return i
			}
		}
	}
}
