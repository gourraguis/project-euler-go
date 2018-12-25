package utils

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

func DivisorsNumber(n int) int {
	res := 1

	for _, exponents := range PrimeFactors(n) {
		res *= exponents + 1
	}

	return res
}
