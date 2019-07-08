package linked_list

import "fmt"

type LinkNode struct {
	next *LinkNode
	val  interface{}
}

func NewLinkNode(val interface{}) *LinkNode {
	return &LinkNode{next: nil, val: val}
}

func (this *LinkNode) GetValue() interface{} {
	return this.val
}

func (this *LinkNode) GetNext() *LinkNode {
	return this.next
}

type LinkedList struct {
	head   *LinkNode
	length uint
}

func NewLinkedList() *LinkedList {
	return &LinkedList{NewLinkNode(struct{}{}), 0}
}

// 在某个节点前插入节点, 返回是否插入成功
func (this *LinkedList) InsertBefore(p *LinkNode, v interface{}) bool {
	if p == nil || p == this.head {
		return false
	}
	prev := this.head
	cur := prev.next
	for cur != nil {
		if cur == p {
			node := NewLinkNode(v)
			node.next = cur
			prev.next = node
			this.length++
			return true
		}
		prev = cur
		cur = cur.next
	}
	return false
}

// 在某个节点后面插入节点, 返回是否插入成功
func (this *LinkedList) InsertAfter(p *LinkNode, v interface{}) bool {
	if p == nil {
		return false
	}
	node := NewLinkNode(v)
	node.next = p.next
	p.next = node
	this.length++
	return true
}

func (this *LinkedList) InsertToHead(v interface{}) bool {
	return this.InsertAfter(this.head, v)
}

func (this *LinkedList) InsertToTail(v interface{}) bool {
	cur := this.head
	for cur != nil && cur.next != nil {
		cur = cur.next
	}
	cur.next = NewLinkNode(v)
	this.length++
	return true
}

func (this *LinkedList) FindByIndex(i uint) *LinkNode {
	if i >= this.length {
		return nil
	}
	var curIdx uint = 0
	cur := this.head.next
	for ; curIdx < i; curIdx++ {
		cur = cur.next
	}
	return cur
}

func (this *LinkedList) DeleteNode(p *LinkNode) bool {
	prev := this.head
	cur := prev.next
	for cur != nil {
		if cur == p {
			prev.next = cur.next
			p = nil
			this.length--
			return true
		}
		prev = cur
		cur = cur.next
	}
	return false
}

// 链表反转
func (this *LinkedList) Reverse() {
	var prev *LinkNode = nil
	cur := this.head.next
	for cur != nil {
		temp := cur.next
		cur.next = prev
		prev = cur
		cur = temp
	}
	this.head.next = prev
}

func (this *LinkedList) String() string {
	s := ""
	cur := this.head.next
	for cur != nil {
		s += fmt.Sprintf("%v", cur.val)
		cur = cur.next
		if cur != nil {
			s += "=>"
		}
	}
	return s
}
