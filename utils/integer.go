package utils

func MaxInt(a, b int64) int64 {
	if a > b {
		return a
	} else {
		return b
	}
}
func MinInt(a, b int64) int64 {
	if a > b {
		return b
	} else {
		return a
	}
}

// 一个整形数组中的最小值和最大值的
func ExtremumInArray(array []int64) (int64, int64) {
	if len(array) < 1 {
		return 0, 0
	}
	min := array[0]
	max := array[0]
	for _, v := range array {
		if v < min {
			min = v
		} else if v > max {
			max = v
		}
	}
	return min, max
}
func MaxIntInSlice(array []int64) int64 {
	_, max := ExtremumInArray(array)
	return max
}
func MinIntInSlice(array []int64) int64 {
	min, _ := ExtremumInArray(array)
	return min
}
