package main

import (
	"fmt"
	"strconv"
	"time"
)

var testMap map[int]*string

func copyBak() map[int]*string {
	vals := make(map[int]*string)

	for key, value := range testMap {
		vals[key] = value
	}

	return vals
}

func main() {

	testMap = make(map[int]*string)
	str := "123"
	testMap[1] = &str
	testMap[2] = &str
	testMap[3] = &str
	testMap[4] = &str

	go func() {
		for {
			fmt.Println(*testMap[3])
			time.Sleep(time.Second)
		}
	}()

	go func() {
		for i := 0; i < 100000000; i++ {
			// 写数据是需要加锁的，并发写的情况下，可能会出现copy多份数据的情况
			// 但此处只有一个协程做此操作
			// 1、复制出新的Map
			newVals := copyBak()
			// 2、修改已有的元素，或添加新元素
			str = "12343264364634" + strconv.Itoa(i)
			newVals[3] = &str
			// 3、将原有的Map地址指向新的Map
			testMap = newVals
		}
	}()

	select {}
}
