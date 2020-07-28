package _interface

import (
	"fmt"
	"io"
	"os"
	"reflect"
	"testing"
	"unsafe"
)

type Person interface {
	Say()
	GetName() (name string)
}
type Lzh struct {
	Name string
}

func (l Lzh) Say() {
	fmt.Println(l.Name)
}
func (l *Lzh) SetName(name string) {
	l.Name = name
}
func (l Lzh) GetName() string {
	fmt.Println(1)
	return "1"
}
func NewPerson() Person {
	return Lzh{
		Name: "lzh",
	}
}

type Dog struct {
	name string
}

func TestInterfaceImplement(t *testing.T) {
	var lzh *Lzh
	var person Person = lzh
	var person1 Person
	fmt.Println(person1 == nil)
	fmt.Println(lzh == nil)
	fmt.Println(person == nil)
	NewPerson().Say()
	dog := Dog{"little pig"}
	dogP := &dog
	dogPtr := uintptr(unsafe.Pointer(dogP))
	namePtr := dogPtr + unsafe.Offsetof(dogP.name)
	fmt.Println("namePtr", namePtr)
}

func TestTypeAssertion(t *testing.T) {
	var r io.Reader
	var w io.Writer
	tty, err := os.OpenFile("/dev/tty", os.O_RDWR, 0)
	if err != nil {
		return
	}
	r = tty
	w = r.(io.Writer)
	fmt.Println(w)
}

// 反射第三定律: 只有可以修改的值才能修改 反射对象可修改，value值必须是可设置的
func TestSet(t *testing.T) {
	var i int64 = 1
	v := reflect.ValueOf(&i)
	v.Elem().SetInt(2)
	fmt.Println(i)
}

func TestName(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	defer func() {
		defer func() {
			panic("panic again and again")
		}()
		panic("panic again")
	}()

	panic("panic once")
}
