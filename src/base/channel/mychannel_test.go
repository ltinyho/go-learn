package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestNewChannel(t *testing.T) {
	size := 10
	c := NewChannel(size)
	var wg sync.WaitGroup
	wg.Add(1)
	go func(idx int) {
		defer wg.Done()
		for {
			c.Push(idx)
			fmt.Println("push")
		}
	}(1)
	go func(idx int) {
		wg.Add(1)
		defer wg.Done()
		for {
			time.Sleep(time.Second)
			val := c.Pop()
			fmt.Println(val)
		}
	}(1)
	wg.Wait()
}
