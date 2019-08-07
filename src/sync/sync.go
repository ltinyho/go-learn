package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var a string
var once sync.Once

func setup() {
	a = "hello, world"
	fmt.Println(1)
}

func doprint() {
	once.Do(setup)
	print(a)
}

func twoprint() {
	go doprint()
	go doprint()
}
func main() {
	twoprint()
	time.Sleep(time.Second)
}

// 只能执行一次
type Once struct {
	m    sync.Mutex
	done uint32
}

func (o *Once) Do(f func()) {
	if atomic.LoadUint32(&o.done) == 1 {
		fmt.Println("once")
		return
	}

	o.m.Lock()
	defer o.m.Unlock()
	if o.done == 0 {
		defer atomic.StoreUint32(&o.done, 1)
		f()
	}
}

var count = 0

func add() {
	once.Do(func() {
		count++
	})
}
