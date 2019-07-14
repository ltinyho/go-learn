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
