package _func

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
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

func f() *int {
	var tmp = 1
	fmt.Println(&tmp)
	return &tmp
}
func g() int {
	x := new(int)
	fmt.Println(x)
	return *x
}

func TestF(t *testing.T) {
	res := f()
	fmt.Println(res)
}

type MyFile struct {
}

func NewMyFile() *MyFile {
	return &MyFile{}
}

func (f *MyFile) Close() {

}

func (f *MyFile) Open() {

}

var CloseFile = (*MyFile).Close

func TestName(t *testing.T) {
	f := NewMyFile()
	CloseFile(f)
}

type UpperWriter struct {
	io.Writer
}

func (p UpperWriter) Write(data []byte) (n int, err error) {
	return p.Writer.Write(bytes.ToUpper(data))
}

func TestUpperWriter(t *testing.T) {
	fmt.Fprintln(&UpperWriter{os.Stdout}, "hello, world")
}

type UppperString string

func (s UppperString) String() string {
	return strings.ToUpper(string(s))
}

func TestUpperString(t *testing.T) {
	fmt.Fprintln(os.Stdout, UppperString("hello,world"))
}

type TB struct {
	testing.TB
}

func (p *TB) Fatal(args ...interface{}) {
	fmt.Println("TB.Fatal disabled!")
}

func TestTb(t *testing.T) {
	var tb testing.TB = new(TB)
	tb.Fatal("f")
	tb1 := TB{}
	tb1.Fatal("f")
}

type names []int

func (n names) change() {
	n[0], n[1] = n[1], n[0]
}
func TestAddName(t *testing.T) {
	n := names{1, 2}
	n.change()
	fmt.Println(n)
}

type persion struct {
	name []string
}

func (p persion) add() {
	p.name = append(p.name, "haha")
}
func TestPersionAdd(t *testing.T) {
	p := persion{}
	p.add()
	fmt.Println(p.name)
}
