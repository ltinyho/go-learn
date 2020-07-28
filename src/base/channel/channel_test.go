package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

// 单向通道,一般用来限制函数声明,约束其他代码的行为
func TestSingleChannel(t *testing.T) {
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

// 测试 chanel copy
func TestChannelCopy(t *testing.T) {
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

func TestArrayCopy(t *testing.T) {
	a1 := [3]int{1, 2, 3}
	b1 := a1
	fmt.Println(a1 == b1) // true
	b1[0] = 2
	fmt.Println(a1[0], b1[0]) // 1 2
	fmt.Println(a1 == b1)     // false
}

func TestChannelSelect(t *testing.T) {
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
}

func TestChannelLimit(t *testing.T) {
	ch1 := make(chan struct{}, 100)
	i := 0
	rand.Seed(time.Now().Unix())
	go func() {
		for {
			time.Sleep(time.Second)
			for j := 0; j < 5; j++ {
				go func() {
					sec := rand.Intn(10) + 5
					time.Sleep(time.Second * time.Duration(sec))
					fmt.Println("sec", sec)
					<-ch1
				}()
			}
		}
	}()
	for {
		ch1 <- struct{}{}
		i++
		fmt.Println("i", i)
	}
}

var done1 = make(chan bool, 2)
var msg string

func aGoroutine() {
	msg = "hello, world"
	done1 <- true
}

// 内存一致性模型
func TestMem1(t *testing.T) {
	go aGoroutine()
	<-done1
	println(msg)
}

var done2 = make(chan bool, 0)

func bGoroutine() {
	msg = "hello, world"
	<-done2
}

// 内存一致性模型
func TestMem2(t *testing.T) {
	go bGoroutine()
	done2 <- true
	println(msg)
}

type line struct {
	l      sync.RWMutex
	online bool
}

func (o *line) Online() {
	o.l.Lock()
	defer o.l.Unlock()
	o.online = true
}

func (o *line) Offline() {
	o.l.Lock()
	defer o.l.Unlock()
	o.online = false
}
func (o *line) isOnline() bool {
	o.l.RLock()
	defer o.l.RUnlock()
	return o.online
}

// 超时控制
func TestTimeOut(t *testing.T) {
	closeCh := time.Tick(time.Second * 2)
	testCh := make(chan int, 10)
	rand.Seed(time.Now().Unix())
	var heart int64 = 1
	_line := &line{}
	var count = 3
	var _count = count
	var reconnCount = 3
	go func() {
		for range time.Tick(time.Second * time.Duration(heart)) {
			val := rand.Intn(10)
			if val > 5 {
				fmt.Println("Offline")
				_line.Offline()
			} else {
				fmt.Println("Online")
				_line.Online()
			}
		}
	}()
test:
	for {
		select {
		case <-closeCh:
			isOnline := _line.isOnline()
			if !isOnline {
				_count--
				if _count <= 0 {
					_count = count
					fmt.Println("reconnection")
					reconnCount--
					if reconnCount <= 0 {
						fmt.Println("too many reconnection")
						break test
					}
				}
			}
		case val := <-testCh:
			fmt.Println(val)
		}
	}
	fmt.Println("end")
}

// 解耦生产者和消费者
func TestConsumer(t *testing.T) {
	tasks := make(chan int, 1)
	workCount := 5
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	closeChan := make(chan int, 0)
	for i := 0; i < workCount; i++ {
		go work(ctx, i, tasks)
	}
	rand.Seed(time.Now().UnixNano())
	go func() {
		for {
			select {
			case <-closeChan:
				cancel()
				fmt.Println("close")
				return
			case <-ctx.Done():
				return
			case tasks <- rand.Intn(1000):
			}
		}
	}()
	select {
	case <-ctx.Done():
		close(tasks)
	case <-time.After(time.Second * 5):
		close(closeChan)
	}
}

func work(ctx context.Context, id int, tasks <-chan int) {
	for {
		select {
		case task := <-tasks:
			fmt.Printf("id:%d,task:%d\n", id, task)
		case <-ctx.Done():
			fmt.Println("work done:", id)
			return
		}
	}
}

func TestClosedChannel(t *testing.T) {
	ch := make(chan int, 10000)
	go func() {
		for {
			if len(ch) == 10000 {
				fmt.Println("closed")
				close(ch)
				break
			}
			ch <- 1
		}
	}()
	go func() {
		var count = 1
		time.Sleep(time.Second * 5)
		for {
			count++
			val := <-ch
			if count <= 10000 {
				fmt.Println(count, val)
			}
		}
	}()
	select {}
}

func TestName(t *testing.T) {
	val := foo()
	t.Log(val)
}
func foo() (i int) {
	defer func() {
		i++
	}()

	return i
}
