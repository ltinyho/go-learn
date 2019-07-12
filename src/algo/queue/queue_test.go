package queue

import (
	"fmt"
	"testing"
)

func TestNewArrayQueue(t *testing.T) {
	queue := NewArrayQueue(4)
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
