package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

// 保证 goroutine 按顺序执行
func main() {
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
