package main

import (
	"fmt"
	"time"
)

func main() {
	for i:=0;i<10;i++{
		defer func(id int) {
			fmt.Println(id)
		}(i)
	}
}

func fetchDemo() {
	defer func() {
		if v := recover(); v != nil {
			fmt.Printf("Recovered a panic. [index=%d]\n", v)
		}
	}()
	ss := []string{"A", "B", "C"}
	fmt.Printf("Fetch the elements in %v one by one...\n", ss)
	fetchElement(ss, 0)
	fmt.Println("The elements fetching is done.")
}

func fetchElement(ss []string, index int) (element string) {
	if index >= len(ss) {
		fmt.Printf("Occur a panic! [index=%d]\n", index)
		panic(index)
	}
	fmt.Printf("Fetching the element ... [index=%d]\n", index)
	element = ss[index]
	defer fmt.Printf("The element is \"%s\". [index=%d]\n", element, index)
	fetchElement(ss, index+1)
	return
}

func runPanic() {
	go func() {
		fmt.Println(IsPanic())
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("Recover success...")
				return
			}
		}()
		panic("ok")
	}()
	for {
		select {
		case <-time.Tick(time.Second):
			fmt.Println("tick")
		}
	}
}
// recover没有被defer方法直接调用,不能捕获panic
func IsPanic() bool {
	if err := recover(); err != nil {
		fmt.Println("Recover success...")
		return true
	}

	return false
}
