package main

import (
	"sync"
	"testing"
	"time"
)

// 内存模型相关例子

func TestWait(t *testing.T) {
	var wg sync.WaitGroup
	var count int
	var ch = make(chan bool, 1)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			ch <- true
			count++
			time.Sleep(time.Millisecond)
			count--
			<-ch
			wg.Done()
		}()
	}
	wg.Wait()
}
func TestIncorrectSync(t *testing.T) {

}
