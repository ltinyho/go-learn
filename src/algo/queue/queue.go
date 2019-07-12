package queue

import "fmt"

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
	// 队列已满
	if q.head == 0 && q.tail == q.limit {
		fmt.Println("队列已满")
		return false
	}
	// 已经到队尾,搬移数据
	if q.tail == q.limit {
		for i := q.head; i < q.tail; i++ {
			q.data[i-q.head] = q.data[i]
			q.data[i] = nil
		}
		q.tail = 0 + (q.tail - q.head)
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
