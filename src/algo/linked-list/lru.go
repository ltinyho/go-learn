package linked_list

import (
	"time"
)

type node struct {
	next  *node
	key   string
	value string
	date  time.Time
}

type LRU struct {
	head    *node
	maxSize int64
	length  int64
}

func NewLRU(maxSize int64) *LRU {
	return &LRU{head: nil, maxSize: maxSize, length: 0}
}

func (l *LRU) add(key, val string) {
	if l.head == nil {
		l.head = newNode(key, val)
		l.length += 1
		return
	}

	if l.length == l.maxSize {
		l.unshift(key, val)
		tailPrev := l.index(l.length - 1)
		tailPrev.next = nil
	} else {
		l.unshift(key, val)
		l.length += 1
	}
}

func (l *LRU) unshift(key, val string) {
	temp := newNode(key, val)
	temp.next = l.head
	l.head = temp
}

func (l *LRU) tail() *node {
	cur := l.head
	for cur != nil && cur.next != nil {
		cur = cur.next
	}
	return cur
}
func (l *LRU) index(i int64) *node {
	var idx int64 = 0
	cur := l.head
	for cur != nil && idx < i {
		idx++
		cur = cur.next
	}
	return cur
}
func newNode(key, val string) *node {
	return &node{
		key:   key,
		value: val,
		date:  time.Now(),
		next:  nil,
	}
}

func (l *LRU) get(key string) (val string) {
	cur := l.head
	prev := &node{}
	prev.next = cur
	for cur != nil {
		if cur.key == key {
			val = cur.value
			prev.next = cur.next
			cur.next = l.head
			l.head = cur
			return
		}
		prev = cur
		cur = cur.next
	}
	return
}
func (l *LRU) delete(val string) {

}

func (l *LRU) size() int64 {
	return l.length
}
func (l *LRU) String() string {
	s := ""
	cur := l.head
	for cur != nil && cur.next != nil {
		s += cur.value + "=>"
		cur = cur.next
	}
	s += cur.value
	return s
}
