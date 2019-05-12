package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("/Volumes/work/go/go-learn/src/file/test.txt")
	defer file.Close()
	if err != nil {
		fmt.Println(err)
	}
	file.WriteString("1")
	file.WriteString("2")
}
