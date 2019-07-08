package linked_list

import (
	"fmt"
	"testing"
)

func TestNewLinkedList(t *testing.T) {
	linkedList := NewLinkedList()
	linkedList.InsertToTail(1)
	fmt.Println(linkedList)
	linkedList.Reverse()
	fmt.Println(linkedList)
}
