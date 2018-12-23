package utils

func SliceSum(slice []int) int {
	res := 0
	for _, elem := range slice {
		res += elem
	}
	return res
}

func SliceProduct(slice []int) int {
	res := 1
	for _, elem := range slice {
		res *= elem
	}
	return res
}

func SliceBiggest(arr []int) int {
	res := arr[0]
	for _, v := range arr {
		if v > res {
			res = v
		}
	}
	return res
}
