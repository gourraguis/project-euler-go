package main

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

//	Problem: What is the sum of the digits of the number 2exp1000?
func main() {
	fmt.Println(powerDigitSum(2, 1000))
}

func powerDigitSum(num, exp int) int {
	res := big.NewInt(1)
	for i := 0; i < exp; i++ {
		res.Mul(res, big.NewInt(int64(num)))
	}
	return digitSum(res.String())
}

func digitSum(num string) int {
	res := 0
	strArr := strings.Split(num, "")
	for _, elem := range strArr {
		convertedNum, err := strconv.Atoi(elem)
		if err != nil {
			panic(err)
		}
		res += convertedNum
	}
	return res
}
