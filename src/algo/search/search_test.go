package search

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	ary1 := []int{0, 1, 2, 3, 3, 4, 5, 6, 7}
	var res int
	res = BinarySearchFirstEq(ary1, 3)
	if res != 3 {
		t.Fatal("第一个等于", res)
	}
	res = BinarySearchFirstGtEq(ary1, 4)
	if res != 5 {
		t.Fatal("第一个大于等于", res)
	}
	res = BinarySearchLastEq(ary1, 3)
	if res != 4 {
		t.Fatal("最后一个等于", res)
	}
	res = BinarySearchLastLtEq(ary1, 6)
	if res != 7 {
		t.Fatal("最后一个小于等于", res)
	}
}

func TestBinarySearchRecursive(t *testing.T) {
	res := BinarySearchRecursive([]int{1, 3, 4, 4, 5, 6}, 3)
	fmt.Println(res)
}
