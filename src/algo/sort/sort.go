package sort

func BubbleSort(array []int) []int {
	res := make([]int, len(array))
	copy(res, array)
	for i := 0; i < len(res); i++ {
		flag := true // 本次排序是否交换过,如果没交换过,说明已经有序了,可以提前退出排序
		for j := 0; j < len(res)-i-1; j++ {
			if res[j] > res[j+1] {
				temp := res[j]
				res[j] = res[j+1]
				res[j+1] = temp
				flag = false
			}
		}

		if flag {
			break
		}
	}
	return res
}

func InsertSort(array []int) []int {
	res := make([]int, len(array))
	copy(res, array)
	for i := 1; i < len(res); i++ {
		value := res[i]
		var j = i
		for ; j > 0; j-- {
			if res[j-1] > value {
				res[j] = res[j-1]
			} else {
				break
			}
		}
		res[j] = value
	}
	return res
}
func SelectSort(array []int) []int {
	res := make([]int, len(array))
	copy(res, array)
	for i := 0; i < len(res); i++ {
		min := i
		for j := i + 1; j < len(res); j++ {
			if res[j] < res[min] {
				min = j
			}
		}
		temp := res[i]
		res[i] = res[min]
		res[min] = temp
	}
	return res
}
func ShellSort(array []int) []int {
	res := make([]int, len(array))
	copy(res, array)
	key := len(res) / 2
	for key > 0 {
		for i := key; i < len(res); i++ {
			for j := i; j >= key && res[j] < res[j-key]; j -= key {
				res[j], res[j-key] = res[j-key], res[j]
			}
		}
		key = key / 2
	}
	return res
}
func MergeSort(array []int) []int {
	length := len(array)
	if length < 2 {
		return array
	}
	mid := length / 2
	left := array[:mid]
	right := array[mid:]
	return MergeArray(MergeSort(left), MergeSort(right))
}

func MergeArray(left []int, right []int) []int {
	var res []int
	i := 0
	j := 0
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			res = append(res, left[i])
			i++
		} else {
			res = append(res, right[j])
			j++
		}
	}
	// 判断哪个子数组中有剩余的数据
	if i == len(left) {
		res = append(res, right[j:]...)
	} else if j == len(right) {
		res = append(res, left[i:]...)
	}
	return res
}

func QuickSort(array []int) []int {
	res := make([]int, len(array))
	copy(res, array)
	quickSort(res, 0, len(res)-1)
	return res
}
func quickSort(array []int, left, right int) {
	if left >= right {
		return
	}
	r := partition(array, left, right)
	quickSort(array, left, r-1)
	quickSort(array, r+1, right)
}
func partition(array []int, left, right int) int {
	pivot := array[right]
	i := left
	for j := left; j < right; j++ {
		if array[j] < pivot {
			array[i], array[j] = array[j], array[i]
			i++
		}
	}
	array[i], array[right] = array[right], array[i]
	return i
}
func qsort(data []int) {
	if len(data) <= 1 {
		return
	}
	mid := data[0]
	head, tail := 0, len(data)-1
	for i := 1; i <= tail; {
		if data[i] > mid {
			data[i], data[tail] = data[tail], data[i]
			tail--
		} else {
			data[i], data[head] = data[head], data[i]
			head++
			i++
		}
	}
	qsort(data[:head])
	qsort(data[head+1:])
}
