package main

import (
	"context"
	"fmt"
	"testing"
)

func TestCtxValue(t *testing.T) {
	ctx := context.Background()
	func1(ctx, "lzh", "ok")
}

// 使用 context 传值,只能从上往下传值
func func1(ctx context.Context, key, val string) {
	ctx = context.WithValue(ctx, key, val)
	func2(ctx, key)
}

func func2(ctx context.Context, key string) {
	val := ctx.Value(key).(string)
	fmt.Println(val)
}

func TestMyCtxValue(t *testing.T) {
	ctx := Background()
	myfunc1(ctx, "lzh", "ok")
}

var myctx = new(valueCtx)

func Background() MyContext {
	return myctx
}

// 使用 context 传值,只能从上往下传值
func myfunc1(ctx MyContext, key, val string) {
	ctx1 := WithValue(ctx, key, val)
	ctx2 := WithValue(ctx1, "haha", "haha")
	myfunc2(ctx2, key)
}

func myfunc2(ctx MyContext, key string) {
	val := ctx.Value(key).(string)
	fmt.Println(val)
}

// 自定义 context
type MyContext interface {
	Value(key string) interface{}
}
type valueCtx struct {
	MyContext
	key, val interface{}
}

func WithValue(ctx MyContext, key, val interface{}) MyContext {
	c := &valueCtx{
		MyContext: ctx,
		key:       key,
		val:       val,
	}
	return c
}
func (c *valueCtx) Value(key string) interface{} {
	if c.key == key {
		return c.val
	}
	return c.MyContext.Value(key)
}

// 使用 context 取消耗时操作



