package main

import (
	"fmt"
	"math/big"
	"project-euler-go/utils"
)

//	Problem: Find the sum of all the primes below two million.
func main() {
	fmt.Println(utils.SliceSum(primesList(2000000)))
}

func primesList(limit int) []int {
	//	we know that 2 is the only even prime number, so put it on the res then skip all other even numbers
	res := []int{2}
	for i := 3; i < limit; i += 2 {
		if big.NewInt(int64(i)).ProbablyPrime(20) {
			res = append(res, i)
		}
	}
	return res
}
