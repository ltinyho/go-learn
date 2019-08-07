package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

// 保证 goroutine 按顺序执行
func main() {
	raceCondition()
}

func goPrint() {
	for {
		go fmt.Println(0)
		fmt.Println(1)
	}
}

func run() {
	var count int64
	num := 100
	trigger := func(i int64, fn func()) {
		for {
			if n := atomic.LoadInt64(&count); n == i {
				// 自旋锁
				fn()
				atomic.AddInt64(&count, 1)
				break
			}
			time.Sleep(time.Nanosecond)
		}
	}
	for i := 0; i < num; i++ {
		go func(idx int64) {
			rand.Seed(time.Now().Unix())
			sec := rand.Intn(20)
			time.Sleep(time.Second * time.Duration(sec))
			trigger(idx, func() {
				fmt.Println(idx)
				fmt.Println("time", idx, sec)
			})
		}(int64(i))
	}
	trigger(int64(num), func() {})
}

var i = 0

func raceCondition() {
	go func() {
		for i < 10 {
			i = i + 1
			fmt.Println("A wins!")
		}
	}()
	go func() {
		for i > -10 {
			i = i - 1
			fmt.Println("B wins!")
		}
	}()
	time.Sleep(time.Hour)
}
