package main

import "fmt"

//	Problem: Which starting number, under one million, produces the longest chain (for Collatz sequence) ?
func main() {
	fmt.Println(longestCollatzChain(1000000))
}

func longestCollatzChain(limit int) int {
	res := 1       //	this is the solution, i.e the starting element
	resLength := 1 // this is the length of the chain using res as starting point

	for i := 1; i < limit; i++ {
		holder := collatzChainLength(i)
		if holder > resLength {
			resLength = holder
			res = i
		}
	}

	return res
}

func collatzChainLength(start int) int {
	res := 1
	for start != 1 {
		res++
		start = nextCollatzElem(start)
	}

	return res
}

func nextCollatzElem(n int) int {
	if n%2 == 0 {
		return n / 2
	} else {
		return 3*n + 1
	}
}
