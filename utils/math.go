package utils

//	Takes a number n and returns a map of primes that divide it and their exponents.
func PrimeFactors(n int) map[int]int {
	res := make(map[int]int)

	for n%2 == 0 {
		if _, ok := res[2]; ok {
			res[2] += 1
		} else {
			res[2] = 1
		}
		n /= 2
	}

	for i := 3; i*i <= n; i += 2 {
		for n%i == 0 {
			if _, ok := res[i]; ok {
				res[i] += 1
			} else {
				res[i] = 1
			}
			n /= i
		}
	}

	if n > 2 {
		res[n] = 1
	}

	return res
}

//	Returns the number of dividers a number has
func DivisorsNumber(n int) int {
	res := 1

	for _, exponents := range PrimeFactors(n) {
		res *= exponents + 1
	}

	return res
}

func Factorial(n int) int {
	if n < 1 {
		return 1
	}

	var res int64 = 1
	for i := 2; i <= n; i++ {
		res *= int64(i)
	}
	return res
}

func Square(n int) int {
	return int(n) * int(n)
}
