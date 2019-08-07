package main

import (
	"fmt"
	"reflect"
	"unicode/utf8"
	"unsafe"
)

func main() {
	s := "hello,world"
	fmt.Println("len(s):", (*reflect.StringHeader)(unsafe.Pointer(&s)).Len)
	fmt.Println("data:", (*reflect.StringHeader)(unsafe.Pointer(&s)).Data)
	s1 := "世界"
	fmt.Printf("%#v , %v \n", []byte(s1), len(s1))
	fmt.Println("\xe4\xb813a13")
	// 将字符串转成[]byte 字节序列,一般不会产生运行时开销
	s2 := "世界abc"
	fmt.Println(utf8.RuneCountInString(s2))
	s3 := []rune(s2)
	fmt.Println(len(s3))
	for k := range s3 {
		fmt.Println(string(s3[k]))
	}
	str := "hello 世界"
	for _, v := range str {
		fmt.Println(string(v))
	}
}
