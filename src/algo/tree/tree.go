package tree

import (
	"fmt"
)

type Node struct {
	data  int
	left  *Node
	right *Node
}

func NewNode(data int) *Node {
	return &Node{data: data}
}
func (this *Node) String() string {
	return fmt.Sprintf("v:%+v, left:%+v, right:%+v", this.data, this.left, this.right)
}

type Tree struct {
	root *Node
}

func NewTree() *Tree {
	return &Tree{root: nil}
}

// 二叉查找树
func (this *Tree) Insert(val int) {
	if this.root == nil {
		this.root = NewNode(val)
		return
	}
	p := this.root
	for p != nil {
		if val > p.data {
			if p.right == nil {
				p.right = NewNode(val)
				return
			}
			p = p.right
		} else {
			if p.left == nil {
				p.left = NewNode(val)
				return
			}
			p = p.left
		}
	}
}

func (this *Tree) Find(val int) *Node {
	p := this.root
	for p != nil {
		if val > p.data {
			p = p.right
		} else if val < p.data {
			p = p.left
		} else {
			return p
		}
	}
	return nil
}

func (this *Tree) Delete(val int) {
	cur := this.root
	var parent *Node
	for cur != nil && cur.data != val {
		parent = cur
		if val > cur.data {
			cur = cur.right
		} else if val < cur.data {
			cur = cur.left
		}
	}
	// 未找到
	if cur == nil {
		return
	}
	//  删除的节点有两个子节点,找到右子树最小节点
	if cur.left != nil && cur.right != nil {
		minCur := cur.right
		minParent := cur
		for minCur.left != nil {
			minParent = minCur
			minCur = minCur.left
		}
		cur.data = minCur.data // 替换要删除的节点数据
		cur = minCur           // 下面变成删除最小的节点,也就是删除叶子节点或者仅有一个节点的操作了
		parent = minParent
	}

	// 删除节点是叶子节点或者仅有一个节点
	var child *Node
	if cur.left != nil {
		child = cur.left
	} else if cur.right != nil {
		child = cur.right
	}
	if parent == nil { // 需要优先判断 parent 是否为空
		this.root = child
	} else if parent.left == cur {
		parent.left = child
	} else if parent.right == cur {
		parent.right = child
	}
}

// 前序遍历
func PrevOrder(root *Node) {
	if root == nil {
		return
	}
	fmt.Printf("%v", root.data)
	PrevOrder(root.left)
	PrevOrder(root.right)
}

// 中序遍历
func InOrder(root *Node) {
	if root == nil {
		return
	}
	InOrder(root.left)
	fmt.Printf("%v", root.data)
	InOrder(root.right)
}

// 后序遍历
func PostOrder(root *Node) {
	if root == nil {
		return
	}
	PostOrder(root.left)
	PostOrder(root.right)
	fmt.Printf("%v", root.data)
}

type Heap struct {
	count int   // 已存数据个数
	data  []int // 数据
	size  int   // 最大容量
}

func NewHeap(size int) *Heap {
	return &Heap{
		count: 0,
		data:  make([]int, size+1),
		size:  size,
	}
}
func (h *Heap) Insert(num int) {
	if h.count >= h.size {
		fmt.Println("full size")
		return
	}
	h.count++
	h.data[h.count] = num
	i := h.count
	for i/2 > 0 && h.data[i] > h.data[i/2] {
		h.data[i], h.data[i/2] = h.data[i/2], h.data[i]
		i = i / 2
	}
}
func (h *Heap) RemoveMax() {
	if h.count <= 0 {
		return
	}
	h.data[1] = h.data[h.count]
	h.data[h.count] = -1
	h.count--
	Heapify(h.data, h.count, 1)
}

func Heapify(data []int, count, i int) {
	for {
		maxPos := i
		if i*2 <= count && data[i] < data[i*2] {
			maxPos = i * 2
		}

		if i*2+1 <= count && data[maxPos] < data[i*2+1] {
			maxPos = i*2 + 1
		}
		if maxPos == i {
			break
		}
		data[i], data[maxPos] = data[maxPos], data[i]
		i = maxPos
	}
}

func BuildHeap(data []int, count int) {
	for i := count / 2; i >= 1; i-- {
		Heapify(data, count, i)
	}
}

func SortHeap(data []int) {
	count := len(data) - 1
	BuildHeap(data, count)
	k := count
	for k > 1 {
		data[k], data[1] = data[1], data[k]
		k--
		Heapify(data, k, 1)
	}
}
