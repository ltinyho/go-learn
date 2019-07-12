package queue

import (
	"fmt"
	"testing"
)

func testQueue(queue Queue) {
	queue.Enqueue(1)
	queue.Dequeue()
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Dequeue()
	queue.Enqueue(5)
	queue.Enqueue(6)
	queue.Dequeue()
	queue.Enqueue(7)
	fmt.Println(queue)
}
func TestNewArrayQueue(t *testing.T) {
	queue := NewArrayQueue(4)
	testQueue(queue)
}

func TestNewLinkedQueue(t *testing.T) {
	queue := NewLinkedQueue()
	testQueue(queue)
}

func TestNewCycleQueue(t *testing.T) {
	queue := NewCycleQueue(4)
	testQueue(queue)
}
