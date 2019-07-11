package stack

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	stack := NewStack()
	var i uint64 = 0
	for i = 0; i < 10000000; i++ {
		stack.Push(i)
	}
	fmt.Println("清除链表")
	stack.Flush()
	fmt.Println(stack)
}

func TestCalExpression(t *testing.T) {
	res := CalExpression("3+5*8-6")
	fmt.Println(res)
}

func TestValidParentheses(t *testing.T) {
	res := ValidParentheses("{}")
	fmt.Println(res)
}
