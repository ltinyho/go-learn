package main

import "fmt"

func main() {
	sliceOperate()
}

func sliceReference() {
	var a = [3]int{1, 2, 3}
	fmt.Println(a[0])
	b := a
	b[1] = 0
	fmt.Println(b[1], a[1])
	c := &b
	c[1] = -1
	fmt.Println(*c == b, a == b)
	fmt.Printf("%T\n", b)
	fmt.Printf("%#v\n", b)
	var s = "kdsfjkls"
	fmt.Printf("%#v\n", s)
}

func sliceOperate() {
	var a []int
	// 尾部添加
	a = append(a, 1)
	a = append(a, 2, 3)
	a = append(a, []int{4, 5, 6}...)
	// 头部添加
	a = append([]int{0}, a...)
	a = append([]int{-3, -2, -1}, a...)
	fmt.Println(a)
	a = append(a[:2], append([]int{1, 2}, a[2:]...)...) // 会生成临时切片
	fmt.Println(a)
	a = append(a, 0)   // 扩展切片空间
	copy(a[3:], a[2:]) // 后移一位
	a[2] = 1           // 插入元素
	fmt.Println(a)
	insertSlice := []int{3, 3}
	a = append(a, insertSlice...)
	copy(a[3+len(insertSlice):], a[2:]) // 后移 len(insertSlice)
	copy(a[3:], insertSlice)            // 插入切片
	fmt.Println(a)
	a = append(a[:3], a[5:]...)
	fmt.Println(a)
	a = a[2:]
	fmt.Println(a)
	a = a[:len(a)-2]
	fmt.Println(a)
	a = a[:2+copy(a[2:], a[3:])]
	fmt.Println(a)
	fmt.Println(append(a, 4))
	fmt.Println(a)
	fmt.Printf("%p\n", a)
	b := &a
	fmt.Printf("%p\n", b)
	c := new(int)
	*c = 1
	d := c
	fmt.Println(c, d)
	*c = 22
	fmt.Println(c, d)
	var i *int
	i = new(int)
	*i = 10
	fmt.Println(*i)
}

func referenceSlice() []int {
	a := []int{1, 2, 3, 4}
	return a[:2]
}
func a() func(d int) int {
	var y = 3
	return func(x int) int {
		return y + x
	}
}
