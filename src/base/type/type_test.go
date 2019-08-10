package _type

import (
	"errors"
	"fmt"
	"testing"
)

type Computer interface {
	Cputype() string
}
type Laptop struct {
	cpuType string
}

func (l Laptop) Cputype() string {
	return l.Cputype()
}

func TestType(t *testing.T) {
	res := Computer(Laptop{cpuType: "i5"})
	fmt.Println(res)
}

func TestTypeString(t *testing.T) {
	res := string([]byte{'a'})
	fmt.Println(res)
	res1 := []rune("golang")
	fmt.Println(res1)
	b := -12312 * 123
	fmt.Println(int8(b))
	var ch chan int
	close(ch)
}
func TestPanic(t *testing.T) {
	defer fmt.Println("done1")
	defer fmt.Println("done2")
	defer fmt.Println("done3")
	defer func() {
		res := recover()
		fmt.Println(res)
	}()
	if true {
		panic(errors.New("123"))
	}
}
