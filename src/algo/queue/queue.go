package queue

import (
	"fmt"
)

type ArrayQueue struct {
	data  []interface{}
	head  int
	tail  int
	limit int
}

func NewArrayQueue(maxLength int) *ArrayQueue {
	return &ArrayQueue{limit: maxLength, data: make([]interface{}, maxLength)}
}

func (q *ArrayQueue) Enqueue(v interface{}) bool {
	// 已经到队尾,搬移数据
	if q.tail == q.limit {
		// 队列已满
		if q.head == 0 {
			fmt.Println("队列已满")
			return false
		}
		for i := q.head; i < q.tail; i++ {
			q.data[i-q.head] = q.data[i]
			q.data[i] = nil
		}
		q.tail -= q.head
		q.head = 0
	}
	q.data[q.tail] = v
	q.tail++
	return true
}
func (q *ArrayQueue) Dequeue() interface{} {
	if q.tail == q.head {
		fmt.Println("队列已经为空")
		return nil
	}
	res := q.data[q.head]
	q.data[q.head] = nil
	q.head++
	return res
}

func (q *ArrayQueue) String() string {
	s := ""
	fmt.Println(q.data)
	for i := q.head; i < q.tail; i++ {
		s += fmt.Sprintf("%v ", q.data[i])
	}
	return s
}

type QueueNode struct {
	next *QueueNode
	val  interface{}
}

func NewQueueNode(val interface{}) *QueueNode {
	return &QueueNode{val: val}
}

type LinkedQueue struct {
	head *QueueNode
	tail *QueueNode
}

func NewLinkedQueue() *LinkedQueue {
	return &LinkedQueue{nil, nil}
}

func (q *LinkedQueue) Enqueue(val interface{}) bool {
	newNode := NewQueueNode(val)
	if q.head == nil {
		q.head = newNode
		q.tail = newNode
		return true
	}
	q.tail.next = newNode
	q.tail = q.tail.next
	return true
}
func (q *LinkedQueue) Dequeue() (val interface{}) {
	if q.head == nil {
		fmt.Println("队列为空")
		return nil
	}
	temp := q.head
	q.head = q.head.next
	return temp.val
}
func (q *LinkedQueue) String() string {
	s := ""
	cur := q.head
	for cur != nil && cur.next != nil {
		s += fmt.Sprintf("%v => ", cur.val)
		cur = cur.next
	}
	s += fmt.Sprintf("%v", cur.val)
	return s
}
