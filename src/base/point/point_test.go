package point

import (
	"fmt"
	"testing"
	"unsafe"
)

type Persion struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

const (
	u, u1, u2 = 3.3, false, "c"
	u4, u5, u6
)

type Ia struct {
	Name string
	Age  int
	ints []int
}

func TestChangePoint(t *testing.T) {
	pp := &Persion{"robert", 11}
	var puptr = uintptr(unsafe.Pointer(pp))
	var npp uintptr = puptr + unsafe.Offsetof(pp.Age)
	fmt.Println(puptr, npp)
}
