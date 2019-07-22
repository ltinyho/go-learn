package _func

import (
	"fmt"
	"testing"
)

var Swap = func(a, b int) (int, int) {
	return b, a
}

func Print(a ...interface{}) {
	fmt.Println(a...)
}

// 闭包对捕获的外部变量并不是传值的方式访问,而是以引用的方式访问.
func Inc() (v int) {
	defer func() { v++ }()
	return 30
}
func TestPrint(t *testing.T) {
	a := []interface{}{1, 2, 3, 4}
	Print(a...) // 1 2 3 4
	Print(a)    // [1 2 3 4]
}

func TestInc(t *testing.T) {
	res := Inc()
	fmt.Println(res)
}

func TestLoopClosure(t *testing.T) {
	for i := 0; i < 4; i++ {
		i := i // 为闭包函数生成独有的变量
		func() {
			fmt.Println(i)
		}()
	}
	for i := 0; i < 4; i++ {
		i := i
		func(id int) {
			fmt.Println(id)
		}(i)
	}
}

// 隐式传递指针变量
func twice(x []int) {
	for i := range x {
		x[i] *= 2
	}
}

func TestSliceChange(t *testing.T) {
	data := []int{1, 2, 3, 4}
	twice(data)
	fmt.Println(data)
}
