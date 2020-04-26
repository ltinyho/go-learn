package main

import (
	"context"
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	time.AfterFunc(time.Second, func() {
		fmt.Println("end")
	})
	time.Sleep(time.Hour)
}

func num() {
	ctx, cancel := context.WithCancel(context.Background())
	var total = 12
	var num int64
	for i := 0; i < total; i++ {
		go addNum(&num, func() {
			if atomic.LoadInt64(&num) == int64(total) {
				cancel()
			}
		})
	}
	<-ctx.Done()
	fmt.Println("End")
}
func addNum(num *int64, fn func()) {
	atomic.AddInt64(num, 1)
	fmt.Println(*num)
	fn()
}
