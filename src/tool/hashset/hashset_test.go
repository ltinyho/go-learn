package hashset

import (
	"fmt"
	"testing"
)

func TestHashSet(t *testing.T) {
	set := NewHashSet()
	set.Add(1)
	set.Add(2)
	set.Add(3)
	other := NewHashSet()
	other.Add(1)
	other.Add(5)
	other.Add(3)
	other.Add(4)
	res := NewSimpleSet()
	fmt.Println(Intersect(res, set, other))
}
