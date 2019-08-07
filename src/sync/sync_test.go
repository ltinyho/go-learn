package main

import (
	"fmt"
	"sync"
	"testing"
)

func TestOnce(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			add()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(count)
}


func TestSetup(t *testing.T) {

}
