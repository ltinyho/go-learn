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
