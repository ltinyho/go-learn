package _interface

import (
	"fmt"
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
	*namePtr = 1
	nameP := (*string)(unsafe.Pointer(namePtr))

}
