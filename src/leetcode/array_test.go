package leetcode

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {
	res := TwoSum([]int{2, 7, 11, 15}, 9)
	fmt.Println(res)
}

func TestRemoveDuplicates1(t *testing.T) {
	nums := []int{0, 1, 1, 2, 2, 3, 4}
	length := RemoveDuplicates1(nums)
	fmt.Println(nums, length)
}
func TestRemoveDuplicates2(t *testing.T) {
	nums := []int{0, 1, 1, 2, 2, 3, 4}
	length := RemoveDuplicates2(nums)
	fmt.Println(nums, length)
}
func TestRemoveElement(t *testing.T) {
	nums := []int{0, 1, 1, 2, 2, 3, 4}
	length := RemoveElement(nums, 2)
	fmt.Println(nums, length)
	for i := 0; i < length; i++ {
		fmt.Println(nums[i])
	}
}

func TestSearchInsert1(t *testing.T) {
	res := SearchInsert1([]int{3, 4, 5, 6, 7, 8, 9}, 5)
	fmt.Println(res)
}
func TestMaxSubArray(t *testing.T) {
	max := MaxSubArray([]int{10, -2, -3, 5,})
	fmt.Println(max)
}

func TestFib(t *testing.T) {
	res := Fib(50)
	fmt.Println(res)
}
func TestFibDp(t *testing.T) {
	res := FibDp(50)
	fmt.Println(res)
}

func TestMergeArray(t *testing.T) {
	arr1 := []int{1, 3, 5}
	arr2 := []int{2, 4, 6}
	res := mergeArray(arr1, arr2, len(arr1), len(arr2))
	fmt.Println(res)
}
func TestContainsNearbyDuplicate(t *testing.T) {
	res := ContainsNearbyDuplicate([]int{1, 0, 1, 1}, 1)
	if !res {
		t.Failed()
	}
}

func TestInvertNum(t *testing.T) {
	res := flipNums([]int{1, 0, 1})
	fmt.Println(res)
}

func TestFlipAndInvertImage(t *testing.T) {
	res := FlipAndInvertImage([][]int{
		{1, 1, 0},
		{1, 0, 1},
		{0, 0, 0},
	})
	fmt.Println(res)
}
