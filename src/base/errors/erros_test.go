package main

import (
	"errors"
	"fmt"
	"os"
	"testing"
)

func TestError(t *testing.T) {
	fmt.Println()
	err := errors.New("haha")
	switch err := err.(type) {
	case CacheError:
		fmt.Println(err.Error())
	}
	os.IsExist(err)
}

func TestDefer(t *testing.T) {
	defer func() {
		if p := recover(); p != nil {
			fmt.Println("panic", p)
		}
	}()
	var a uint
	a = 0
	a = ^a
	fmt.Println(a)
}
