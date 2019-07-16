package search

func BinarySearch(array []int, num int) bool {
	low := 0
	high := len(array) - 1
	for low <= high {
		mid := low + (high-low)>>1
		if array[mid] == num {
			return true
		} else if array[mid] > num {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return false
}

// 查找第一个等于 num的值
func BinarySearchFirstEq(array []int, num int) int {
	low := 0
	high := len(array) - 1
	for low <= high {
		mid := low + (high-low)>>1
		if array[mid] == num {
			if mid == 0 || array[mid-1] != num {
				return mid
			} else {
				high = mid - 1
			}
		} else if array[mid] > num {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

// 查找最后一个等于 num的值
func BinarySearchLastEq(array []int, num int) int {
	low := 0
	high := len(array) - 1
	for low <= high {
		mid := low + (high-low)>>1
		if array[mid] == num {
			if mid == len(array)-1 || array[mid+1] != num {
				return mid
			} else {
				low = mid + 1
			}
		} else if array[mid] > num {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

// 查找第一个大于等于给定值的元素
func BinarySearchFirstGtEq(array []int, num int) int {
	low := 0
	high := len(array) - 1
	for low <= high {
		mid := low + (high-low)>>1
		if array[mid] >= num {
			if mid == 0 || array[mid-1] < num {
				return mid
			} else {
				high = mid - 1
			}
		} else {
			low = mid + 1
		}
	}
	return -1
}

// 查找第一个小于等于给定值的元素
func BinarySearchLastLtEq(array []int, num int) int {
	low := 0
	high := len(array) - 1
	for low <= high {
		mid := low + (high-low)>>1
		if array[mid] <= num {
			if mid == len(array)-1 || array[mid+1] > num {
				return mid
			} else {
				low = mid + 1
			}
		} else {
			high = mid - 1
		}
	}
	return -1
}

func BinarySearchRecursive(array []int, num int) bool {
	return binarySearchRecursive(array, 0, len(array)-1, num)
}

func binarySearchRecursive(array []int, low, high, num int) bool {
	if low > high {
		return false
	}
	mid := low + (high-low)>>1
	if array[mid] == num {
		return true
	} else if array[mid] > num {
		return binarySearchRecursive(array, low, mid-1, num)
	} else {
		return binarySearchRecursive(array, mid+1, high, num)
	}
}
