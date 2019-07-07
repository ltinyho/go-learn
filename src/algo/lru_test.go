package algo

import (
	"fmt"
	"testing"
)

func TestLru(t *testing.T) {
	l := NewLRU(5)
	l.add("1", "1")
	l.add("2", "2")
	l.add("3", "3")
	l.add("4", "4")
	l.add("6", "6")
	l.add("7", "7")
	l.get("4")
	l.get("2")
	fmt.Println(l)
}
