package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Book struct {
	sync.RWMutex
	GoBook map[int]int
}

func main() {
	lockMap()
}

func lockMap() {
	testMap := map[int]int{}
	book := Book{
		GoBook: make(map[int]int),
	}
	go func() {
		for {
			book.Lock()
			testMap[1] = rand.Intn(100)
			book.Unlock()
		}
	}()
	go func() {
		for {
			book.RLock()
			fmt.Println(testMap[1])
			book.RUnlock()
		}
	}()
	time.Sleep(time.Second)
}
