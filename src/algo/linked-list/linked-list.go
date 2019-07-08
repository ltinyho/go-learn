package linked_list

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
	return true
}

// 在某个节点后面插入节点, 返回是否插入成功
func (this *LinkedList) InsertAfter(p *LinkNode, v interface{}) bool {
	return true
}

func (this *LinkedList) InsertToHead(v interface{}) bool {
	return true
}

func (this *LinkedList) InsertToTail(v interface{}) bool {
	return true
}

func (this *LinkedList) FindByIndex(i uint) *LinkNode {
	return nil
}

func (this *LinkedList) DeleteNode(p *LinkNode) bool {
	return true
}

func (this *LinkedList) String() string {
	s := ""
	return s
}
