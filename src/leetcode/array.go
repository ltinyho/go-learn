package leetcode

import (
	"fmt"
)

func TwoSum(nums []int, target int) []int {
	var notMap = map[int]int{}
	for i, num := range nums {
		res := target - num
		val, ok := notMap[res]
		if ok {
			return []int{val, i}
		} else {
			notMap[num] = i
		}
	}
	return nil
}

func RemoveDuplicates1(nums []int) int {
	count := 0
	n := len(nums)
	for i := 1; i < n; i++ {
		if nums[i] == nums[i-1] {
			count++
		} else {
			nums[i-count] = nums[i]
		}
	}
	return n - count
}
func RemoveDuplicates2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}
func RemoveElement(nums []int, val int) int {
	i := 0
	for j := 0; j < len(nums); j++ {
		if nums[j] != val {
			nums[i] = nums[j]
			i += 1
		}
	}
	return i
}

func SearchInsert(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		if target <= nums[i] {
			return i
		}
	}
	return len(nums)
}
func SearchInsert1(nums []int, target int) int {
	left := 0
	right := len(nums)
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] >= target {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

// MaxSubArray(nums []int,i,j int)   MaxSubArray(nums,0,1) 改为
// maxSubArray(A, i) 为 maxSubArray A[0:i]
// maxSubArray(A, i) = maxSubArray(A, i - 1) > 0 ? maxSubArray(A, i - 1) : 0 + A[i];
// 1,-2,-5,1,-2
// dp -1,-2,-3,5,6,7,8,9,-1,-1,-2
func MaxSubArray(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	dp := make([]int, n)
	dp[0] = nums[0]
	max := dp[0]
	left := 0
	right := 0
	for i := 1; i < n; i++ {
		if dp[i-1] > 0 {
			dp[i] = nums[i] + dp[i-1]
		} else {
			dp[i] = nums[i] + 0
			left = i
		}
		if max < dp[i] {
			right = i
			max = dp[i]
		}
	}
	fmt.Println(nums[left:right+1], dp)
	return max
}

// 菲波那切数列 递归版
func Fib(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	return Fib(n-1) + Fib(n-2)
}

// 菲波那切数列 动态规划
func FibDp(n int) uint64 {
	mem := make([]uint64, n)
	mem[0] = 1
	mem[1] = 1
	for i := 2; i < n; i++ {
		mem[i] = mem[i-1] + mem[i-2]
	}
	fmt.Println(mem)
	return mem[n-1]
}
func PlusOne(nums []int) []int {
	plus := 1
	length := len(nums)
	for i := length - 1; i > 0; i-- {
		if i == length-1 && nums[i] != 9 {
			nums[i] += 1
			return nums
		}

		if nums[i] == 9 {
			plus = 1
			nums[i] = 0
		} else {
			if plus > 0 {
				nums[i] = nums[i] + 1
				plus = 0
			}
		}
	}
	if plus > 0 {
		nums = append([]int{1}, nums...)
	}

	return nums
}

// 有序数组合并
func mergeArray(nums1, nums2 []int, m, n int) []int {
	//扩容
	nums1 = append(nums1, make([]int, n)...)
	for n > 0 {
		if m > 0 && nums1[m-1] > nums2[n-1] {
			//
			nums1[n+m-1] = nums1[m-1]
			m--
		} else {
			nums1[n+m-1] = nums2[n-1]
			n--
		}
	}
	return nums1
}
func ContainsNearbyDuplicate(nums []int, k int) bool {
	existMap := map[int]int{}
	for key, val := range nums {
		mapVal, ok := existMap[val]
		if ok && (key-mapVal) <= k {
			return true
		}

		existMap[val] = key
	}
	return false
}
