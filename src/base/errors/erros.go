package main

import "fmt"

func main() {
	fetchDemo()
	fmt.Println("The main function is executed.")
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
