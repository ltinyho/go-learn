package search

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	res := BinarySearch([]int{1}, 1)
	fmt.Println(res)
}

func TestBinarySearchRecursive(t *testing.T) {
	res := BinarySearchRecursive([]int{1, 2, 3, 4, 5, 6}, 5)
	fmt.Println(res)
}
