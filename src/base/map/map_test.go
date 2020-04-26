package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestConCurrentMap(t *testing.T) {
	_map := map[string]int{}
	var  _map1 sync.Map
	for i := 0; i < 100; i++ {
		go func() {
				_map["1"]= 100
				_map1.Store("1",100)
		}()
		go func() {
			fmt.Println(_map["1"])
			fmt.Println(_map1.Load("1"))
		}()
	}
	time.Sleep(time.Minute)
}
