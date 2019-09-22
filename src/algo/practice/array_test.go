package practice

import (
	"fmt"
	"testing"
)

// 动态数组
type Array struct {
	data []interface{}
	size int
}

func NewArray(capacity int) *Array {
	return &Array{size: 0, data: make([]interface{}, capacity)}
}
func (a *Array) Capacity() int {
	return len(a.data)
}
func (a *Array) Size() int {
	return a.size
}
func (a *Array) Add(pos int, num interface{}) {
	if pos < 0 || pos > a.Capacity() {
		panic("数组越界")
	}
	if a.Size() == a.Capacity() {
		a.resize(a.Capacity() * 2)
	}
	a.data = append(a.data, 0)
	for i := a.Size(); i > pos; i-- {
		a.data[i] = a.data[i-1]
	}
	a.data[pos] = num
	a.size++
}

func (a *Array) resize(size int) {
	newArray := make([]interface{}, size)
	for e := range a.data {
		newArray = append(newArray, a.data[e])
	}
	a.data = newArray
}
func (a *Array) Push(num interface{}) {
	a.Add(a.Size(), num)
}
func (a *Array) Unshift(num interface{}) {
	a.Add(0, num)
}
func (a *Array) Remove(pos int) {
	if pos < 0 || pos > len(a.data) {
		panic("数组越界")
	}
	if a.Size() == a.Capacity()/4 && a.Size()/2 != 0 {
		a.resize(a.Size() / 2)
	}
	for i := pos; i < a.Size(); i++ {
		a.data[i] = a.data[i+1]
	}
	a.data[a.size] = nil
	a.size--
}
func (a *Array) Pop() {
	a.Remove(a.Size() - 1)
}
func (a *Array) String() string {
	return fmt.Sprintf("%v", a.data)
}
func TestArray(t *testing.T) {
	a := NewArray(20)
	for i := 0; i < 10; i++ {
		a.Push(i)
	}
	fmt.Println(a, a.Capacity())
}
