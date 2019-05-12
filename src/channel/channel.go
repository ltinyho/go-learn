package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Person struct {
	Name string
}
type Book struct {
	person Person
	name   string
}

func main() {
	singleChannel()
	testSelect()
	//testArrayCopy()
	//channelCopy()
}

// 单向通道,一般用来限制函数声明,约束其他代码的行为
func singleChannel() {
	ch1 := make(chan<- int, 1) // 发通道,发送到通道
	ch2 := make(<-chan int, 1) // 收通道,从通道接收
	ch3 := make(chan int, 1)
	ch3 <- 1
	<-ch3
	ch1 <- 2
	ch4 := getIntChan()
	b := <-ch4
	fmt.Println(b, ch2)
}
func getIntChan() <-chan int {
	num := 5
	ch := make(chan int, num)
	for i := 0; i < num; i++ {
		ch <- i
	}
	close(ch)
	return ch
}

// 测试 chanel copy
func channelCopy() {
	ch := make(chan []int, 1)
	s1 := []int{1, 2, 3}
	ch <- s1
	s2 := <-ch

	s2[0] = 100
	fmt.Println(s1, s2) //[100 2 3] [100 2 3]

	//
	ch2 := make(chan [3]int, 1)
	s3 := [3]int{1, 2, 3}
	ch2 <- s3
	s4 := <-ch2

	s3[0] = 100
	fmt.Println(s3, s4) //[100 2 3] [1 2 3]
}

func testArrayCopy() {
	a1 := [3]int{1, 2, 3}
	b1 := a1
	fmt.Println(a1 == b1) // true
	b1[0] = 2
	fmt.Println(a1[0], b1[0]) // 1 2
	fmt.Println(a1 == b1)     // false
}

func testSelect() {
	// 准备好几个通道。
	intChannels := [3]chan int{
		make(chan int, 1),
		make(chan int, 1),
		make(chan int, 1),
	}
	// 随机选择一个通道，并向它发送元素值。
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 1000; i++ {
		index := rand.Intn(3)
		fmt.Printf("The index: %d\n", index)
		// 哪一个通道中有可取的元素值，哪个对应的分支就会被执行。
		// TODO 通道没有值
		go func(idx int) {
			intChannels[idx] <- idx
		}(index)
	}
	for i := 0; i < 1000; i++ {
		select {
		case <-intChannels[0]:
			fmt.Println("The first candidate case is selected.")
		case <-intChannels[1]:
			fmt.Println("The second candidate case is selected.")
		case elem := <-intChannels[2]:
			fmt.Printf("The third candidate case is selected, the element is %d.\n", elem)
		}
	}
	time.Sleep(time.Hour)
}
