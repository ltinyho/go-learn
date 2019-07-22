package tree

import "fmt"

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
