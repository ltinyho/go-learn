package recursive

var res = map[int]int{}

func f(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	if res[n] != 0 {
		return res[n]
	}
	ret := f(n-1) + f(n-2)
	res[n] = ret
	return ret
}
