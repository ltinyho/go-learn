package main

import (
	"fmt"
	"os"
	"testing"
)

func TestOpenFile(t *testing.T) {
	file, err := os.Open("./test.txt")
	if err != nil {
		if pe, ok := err.(*os.PathError); ok {
			fmt.Println(pe.Err, pe.Op)
			fmt.Println("文件不存在")
		}
	} else {
		fmt.Println(file.Name())
	}
}
