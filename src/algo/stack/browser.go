package stack

import "fmt"

type Browser struct {
	backStack    *Stack
	forwardStack *Stack
	curPage      interface{}
}

func NewBrowser() *Browser {
	return &Browser{backStack: NewStack(), forwardStack: NewStack()}
}

func (b *Browser) Back() {
	if b.backStack.IsEmpty() {
		return
	}
	top := b.backStack.Pop()
	b.forwardStack.Push(top.val)
	fmt.Printf("back to %+v\n", top.val)
}

func (b *Browser) Forward() {
	if b.forwardStack.IsEmpty() {
		return
	}
	top := b.forwardStack.Pop()
	b.backStack.Push(top.val)
	fmt.Printf("forward to %+v\n", top.val)
}
func (b *Browser) Open(v interface{}) {
	if v != nil {
		b.forwardStack.Flush()
		b.backStack.Push(v)
		b.curPage = v
		fmt.Printf("Open new addr %+v\n", v)
	}
}
func (b *Browser) CanBack() bool {
	return !b.backStack.IsEmpty()
}
func (b *Browser) CanForward() bool {
	return !b.forwardStack.IsEmpty()
}
