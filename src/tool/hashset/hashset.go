package hashset

import (
	"bytes"
	"fmt"
)

type HashSet struct {
	v map[interface{}]bool
}

type Set interface {
	Add(e interface{}) bool
	Remove(e interface{})
	Clear()
	Contains(e interface{}) bool
	Len() int
	Same(other Set) bool
	Elements() []interface{}
	String() string
}

func NewHashSet() *HashSet {
	return &HashSet{
		v: make(map[interface{}]bool),
	}
}

func (h *HashSet) Add(e interface{}) bool {
	if !h.v[e] {
		h.v[e] = true
		return true
	}
	return false
}
func (h *HashSet) Remove(e interface{}) {
	h.v[e] = false
}

func (h *HashSet) Delete(e interface{}) {
	delete(h.v, e)
}
func (h *HashSet) Clear() {
	h.v = make(map[interface{}]bool)
}
func (h *HashSet) Contains(e interface{}) bool {
	return h.v[e]
}
func (h *HashSet) Len() int {
	return len(h.v)
}

func (h *HashSet) Same(other Set) bool {
	if other == nil {
		return false
	}
	if h.Len() != other.Len() {
		return false
	}
	for key := range h.v {
		if !other.Contains(key) {
			return false
		}
	}
	return true
}

func (h *HashSet) Elements() []interface{} {
	initialLen := len(h.v)
	snapshot := make([]interface{}, initialLen)
	actualLen := 0
	for k := range h.v {
		if actualLen < initialLen {
			snapshot[actualLen] = k
		} else {
			snapshot = append(snapshot, k)
		}
		actualLen++
	}
	if actualLen < initialLen {
		snapshot = snapshot[:actualLen]
	}
	return snapshot
}

func (h *HashSet) String() string {
	strBuffer := bytes.Buffer{}
	strBuffer.WriteString("Set{")
	first := true
	for key := range h.v {
		if first {
			first = false
		} else {
			strBuffer.WriteString(" ")
		}
		strBuffer.WriteString(fmt.Sprintf("%v", key))
	}
	strBuffer.WriteString("}")
	return strBuffer.String()
}

// 超集
func IsSuperset(h Set, other Set) bool {
	if other == nil {
		return false
	}
	curLen := h.Len()
	otherLen := other.Len()
	if curLen == 0 || curLen == otherLen {
		return false
	}

	if curLen > 0 && otherLen == 0 {
		return true
	}

	for _, v := range other.Elements() {
		if !h.Contains(v) {
			return false
		}
	}
	return true
}

// 并集
func Union(h Set, other Set) Set {
	curLen := h.Len()
	otherLen := other.Len()
	if curLen == 0 && otherLen == 0 {
		return nil
	}
	newHashset := NewHashSet()
	if h == nil {
		return nil
	}
	for e := range h.Elements() {
		newHashset.Add(e)
	}
	if other != nil {
		for _, v := range other.Elements() {
			if !newHashset.Contains(v) {
				newHashset.Add(v)
			}
		}
	}

	return newHashset
}

// 交集
func Intersect(res, h Set, other Set) Set {
	if other == nil {
		return nil
	}
	curLen := h.Len()
	otherLen := other.Len()
	if curLen == 0 || otherLen == 0 {
		return nil
	}
	for _, k := range h.Elements() {
		if other.Contains(k) {
			res.Add(k)
		}
	}
	return res
}

func NewSimpleSet() Set {
	return NewHashSet()
}

// 差集
func Difference(res, h Set, other Set) Set {
	if other == nil {
		return nil
	}
	curLen := h.Len()
	otherLen := other.Len()
	if curLen == 0 || otherLen == 0 {
		return nil
	}
	interSet := Intersect(res, h, other)
	union := Union(h, other)
	for _, v := range union.Elements() {
		if !interSet.Contains(v) {
			res.Add(v)
		}
	}
	return res
}
