package main

import (
	"fmt"
	"hash/crc32"
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

func TestMap(t *testing.T) {
	ok := []byte{0x01, 0x02, 0x03, 0x04}
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			if idx%2 == 0 {
				ok[2] = 0x03
			} else {
				ok = []byte{0x00, 0x01, 0x02}
				ok = append(ok, byte(crc32.ChecksumIEEE(ok)))
			}
		}(i)
	}
	wg.Wait()
	if byte(crc32.ChecksumIEEE(ok[:3])) != ok[3] {
		fmt.Println("false")
	}
}
