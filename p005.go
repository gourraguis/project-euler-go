package main

import "fmt"

//	Problem: What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
//	We'll go through all the possible numbers until we find the smallest solution, not optimal but numbers are enough for us to not notice the difference
func main() {
	fmt.Println(smallestMultiple(20))
}

func smallestMultiple(limit int) int {
	arr := createSlice(limit)
	res := 0
	for i := 1; res == 0; i++ {
		isSolution := true
		for _, elem := range arr {
			if i%elem != 0 {
				isSolution = false
				break
			}
		}
		if isSolution {
			res = i
		}
	}
	return res
}

func createSlice(limit int) []int {
	var res []int
	for i := 1; i <= limit; i++ {
		res = append(res, i)
	}
	return res
}
