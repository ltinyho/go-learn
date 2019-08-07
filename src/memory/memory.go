package main

import (
	"fmt"
	"sync"
)

var a, b int

func f() {
	a = 1
	b = 2
}

var wg sync.WaitGroup

func g() {
	print(b)
	print(a)
	wg.Done()
}

type Mystring string
type MyStrings [3]string
type Persion struct {
	name string
}

func main() {
	fmt.Println(0123)
}
