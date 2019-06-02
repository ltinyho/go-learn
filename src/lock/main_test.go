package lock

import (
	"fmt"
	"testing"
)

func TestUnReetrantLock(t *testing.T) {
	l := new(MyLock)
	l.Lock()
	fmt.Println(1)
	l.UnLock()
	fmt.Println(4)
}
