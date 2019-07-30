package tree

import (
	"fmt"
	"testing"
)

func TestPrevOrder(t *testing.T) {

}

func TestInsertTree(t *testing.T) {
	tree := NewTree()
	tree.Insert(4)
	tree.Insert(3)
	tree.Insert(5)
	tree.Insert(2)
	tree.Insert(1)
	tree.Insert(6)
	PrevOrder(tree.root)
	fmt.Println("")
	InOrder(tree.root)
	fmt.Println("")
	PostOrder(tree.root)
	fmt.Println("")
	node := tree.Find(4)
	fmt.Println(node.left.data, node.right.data)
	tree.Delete(4)
	tree.Delete(3)
	InOrder(tree.root)
}

func TestNewHeap(t *testing.T) {
	h := NewHeap(10)
	h.Insert(1)
	h.Insert(2)
	h.Insert(3)
	h.Insert(4)
	h.Insert(5)
	h.Insert(6)
	h.Insert(7)
	fmt.Println(h.data)
	h.RemoveMax()
	fmt.Println(h.data)
}

func TestHeapify(t *testing.T) {
	var data = []int{-1, 1, 2, 3, 4, 5, 6, 7, -1, -1, -1}
	BuildHeap(data, 7)
	fmt.Println(data)
}

func TestSortHeap(t *testing.T) {
	var data = []int{-1, 7, 3, 6, 4, 5, 2, 1}
	SortHeap(data)
	fmt.Println(data)
}
