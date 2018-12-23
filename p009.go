package main

import "fmt"

//	Problem: There exists exactly one Pythagorean triplet for which a + b + c = 1000, find the product abc.
func main() {
	fmt.Println(specialPythagoreanTriplet())
}

func specialPythagoreanTriplet() int {
	for a := 0; a < 1000; a++ {
		for b := a + 1; b < 1000; b++ {
			for c := b + 1; c < 1000; c++ {
				if a*a+b*b == c*c && a+b+c == 1000 {
					return a * b * c
				}
			}
		}
	}
	//	Code never reaches this part, return is here to avoid compile error
	return -1
}
