package stack

import (
	"fmt"
)

type StackNode struct {
	next *StackNode
	val  interface{}
}

func NewStackNode(val interface{}) *StackNode {
	return &StackNode{next: nil, val: val}
}

type Stack struct {
	head   *StackNode
	length uint
}

func NewStack() *Stack {
	return &Stack{NewStackNode(0), 0}
}
func (this *Stack) IsEmpty() bool {
	return this.length == 0
}
func (this *Stack) Pop() *StackNode {
	if this.length == 0 {
		return nil
	}
	top := this.head.next
	this.head.next = top.next
	this.length--
	return top
}
func (this *Stack) Push(v interface{}) {
	newNode := NewStackNode(v)
	newNode.next = this.head.next
	this.head.next = newNode
	this.length++
}

func (this *Stack) Top() *StackNode {
	return this.head.next
}

func (this *Stack) String() string {
	s := ""
	cur := this.head.next
	for cur != nil {
		s += fmt.Sprintf("%v", cur.val)
		cur = cur.next
		if cur != nil {
			s += " => "
		}
	}
	return s
}
func (this *Stack) Flush() {
	prev := this.head
	cur := prev.next
	for cur != nil {
		prev.next = nil
		prev = cur
		cur = cur.next
		this.length--
	}
}

var zeroCode = "0"[0]
var nineCode = "9"[0]
var plusCode = "+"[0]
var subCode = "-"[0]
var mulCode = "*"[0]
var divCode = "/"[0]
var operatorMap = map[string]int{
	"+": 1,
	"-": 1,
	"*": 2,
	"/": 2,
}

// 1+13*9+44-12/3
func CalExpression(exp string) uint {
	/*	operatorStack := NewStack()
		numsStack := NewStack()
		left := 0
		res := 0
		for i := 0; i < len(exp); i++ {
			item := exp[i]
			if item < zeroCode || item > nineCode {
				num, err := strconv.Atoi(exp[left:i])
				if err != nil {
					fmt.Println(err)
				}
				top := operatorStack.Top()
				if top != nil && isTopPriority(top.val.(string), item) {
					num1 := numsStack.Pop().val.(int)
					num2 := numsStack.Pop().val.(int)
					res = cal(num1, num2, item)
				} else {
					operatorStack.Push(item)
				}
				left = i + 1
			} else {
				numsStack.Push(item)
			}
		}
		fmt.Println(numsStack)
		fmt.Println(operatorStack)*/
	return 0
}

func isTopPriority(top, cur string) bool {
	topLevel := operatorMap[top]
	curLevel := operatorMap[cur]
	return topLevel >= curLevel
}
func cal(num1, num2 int, operate string) int {
	switch operate {
	case "+":
		return num1 + num2
	case "-":
		return num1 - num2
	case "/":
		return num1 / num2
	case "*":
		return num1 * num2
	}
	return 0
}

var parenthesesMap = map[string]string{
	"(": ")",
	"[": "]",
	"{": "}",
}

func isLeftParentheses(str string) bool {
	return parenthesesMap[str] != ""
}
func ValidParentheses(str string) bool {
	stack := NewStack()
	for i := 0; i < len(str); i++ {
		item := str[i : i+1]
		if isLeftParentheses(item) {
			stack.Push(item)
		} else {
			top := stack.Pop()
			if parenthesesMap[top.val.(string)] != item {
				return false
			}
		}
	}
	return stack.IsEmpty()
}
