package main

import (
	"container/list"
	"sync"
)

// 通过条件变量实现带缓冲 channel

type IChannel interface {
	Push(v interface{})
	Pop() interface{}
	TryPop() (v interface{}, ok bool)
	TryPush(v interface{}) bool
}

type Channel struct {
	size  int
	queue *list.List
	mutex sync.Mutex
	cond  *sync.Cond
}

func NewChannel(size int) *Channel {
	if size<1{
		panic("todo:support unbuffered channel")
	}
	c:=new (Channel)
	c.cond=sync.NewCond(&c.mutex)
	c.queue= list.New()
	c.size=size
	return c
}

func (c *Channel) Push(v interface{}) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	for c.queue.Len() == c.size {
		c.cond.Wait()
	}
	if c.queue.Len() == 0 {
		c.cond.Broadcast()
	}
	c.queue.PushBack(v)
}

func (c *Channel) Pop() interface{} {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	for c.queue.Len() == 0 {
		c.cond.Wait()
	}
	if c.queue.Len() == c.size {
		c.cond.Broadcast()
	}

	return c.queue.Remove(c.queue.Front())
}



func (c *Channel) TryPush(v interface{}) (ok bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if c.queue.Len() == c.size {
		return
	}
	if c.queue.Len() == 0 {
		c.cond.Broadcast()
	}
	c.queue.InsertAfter(v, c.queue.Back())
	return true
}
func (c *Channel) TryPop() (v interface{}, ok bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if c.queue.Len() == 0 {
		return
	}
	if c.queue.Len() == c.size {
		c.cond.Broadcast()
	}
	return c.queue.Front(),true
}
