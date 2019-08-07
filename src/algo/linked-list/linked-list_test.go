package linked_list

import (
	"fmt"
	"testing"
)

func TestNewLinkedList(t *testing.T) {
	linkedList := NewLinkedList()
	linkedList.InsertToTail(1)
	linkedList.InsertToTail(2)
	linkedList.InsertToTail(3)
	linkedList.InsertToHead(4)
	node := linkedList.FindByIndex(2)
	fmt.Println(node)
	linkedList.DeleteNode(node)
	fmt.Println(linkedList)
	linkedList.Reverse()
	fmt.Println(linkedList)
}
