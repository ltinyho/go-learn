package orderedMap

import (
	"bytes"
	"fmt"
	"reflect"
	"sort"
)

type myOrderedMap struct {
	keys        Keys
	m           map[interface{}]interface{}
	elementType reflect.Type
}

func NewOrderedMap(keys Keys, elementType reflect.Type) *myOrderedMap {
	return &myOrderedMap{keys: keys, m: map[interface{}]interface{}{}, elementType: elementType}
}

func (m *myOrderedMap) Get(key interface{}) interface{} {
	return m.m[key]
}
func (m *myOrderedMap) isAcceptableElem(e interface{}) bool {
	if e == nil {
		return false
	}
	if reflect.TypeOf(e) != m.elementType {
		return false
	}
	return true
}
func (m *myOrderedMap) Put(key interface{}, elem interface{}) (interface{}, bool) {
	if !m.isAcceptableElem(elem) {
		return nil, false
	}
	old, ok := m.m[key]
	m.m[key] = elem
	if !ok {
		m.keys.Add(key)
	}
	return old, true
}

func (m *myOrderedMap) Remove(key interface{}) interface{} {
	old, ok := m.m[key]
	delete(m.m, key)
	if ok {
		m.keys.Remove(key)
	}
	return old
}

func (m *myOrderedMap) Clear() {
	m.keys.Clear()
	m.m = map[interface{}]interface{}{}
}

func (m *myOrderedMap) Len() int {
	return m.keys.Len()
}

func (m *myOrderedMap) Contains(key interface{}) bool {
	_, ok := m.m[key]
	return ok
}

func (m *myOrderedMap) FirstKey() interface{} {
	if m.keys.Len() == 0 {
		return nil
	}
	return m.keys.Get(0)
}

func (m *myOrderedMap) LastKey() interface{} {
	length := m.keys.Len()
	if length == 0 {
		return nil
	}
	return m.keys.Get(length - 1)
}

func (m *myOrderedMap) HeadMap(toKey interface{}) OrderedMap {
	return m.SubMap(nil, toKey)
}

func (m *myOrderedMap) SubMap(fromKey interface{}, toKey interface{}) OrderedMap {
	newMap := NewOrderedMap(NewKeys(m.keys.CompareFunc(), m.KeyType()), m.ElemType())
	mLen := m.Len()
	startIndex, contains := m.keys.Search(fromKey)
	if !contains {
		startIndex = 0
	}
	endIndex, contains := m.keys.Search(toKey)
	if !contains {
		endIndex = mLen
	}
	for i := startIndex; i < endIndex; i++ {
		key := m.keys.Get(i)
		newMap.Put(key, m.m[key])
	}
	return newMap
}

func (m *myOrderedMap) TailMap(fromKey interface{}) OrderedMap {
	return m.SubMap(fromKey, nil)
}

func (m *myOrderedMap) Keys() []interface{} {
	return m.keys.GetAll()
}

func (m *myOrderedMap) Elems() []interface{} {
	initialLen := len(m.m)
	snapshot := make([]interface{}, initialLen)
	actualLen := 0
	for _, v := range m.m {
		if actualLen < initialLen {
			snapshot[actualLen] = v
		} else {
			snapshot = append(snapshot, v)
		}
		actualLen++
	}
	if actualLen < initialLen {
		snapshot = snapshot[:actualLen]
	}
	return snapshot
}

func (m *myOrderedMap) ToMap() map[interface{}]interface{} {
	replica := make(map[interface{}]interface{})
	for k, v := range m.m {
		replica[k] = v
	}
	return replica
}

func (m *myOrderedMap) KeyType() reflect.Type {
	return m.keys.ElemType()
}

func (m *myOrderedMap) ElemType() reflect.Type {
	return m.elementType
}
func (m *myOrderedMap) String() string {
	var buf bytes.Buffer
	buf.WriteString("OrderedMap<")
	buf.WriteString(m.keys.ElemType().Kind().String())
	buf.WriteString(",")
	buf.WriteString(m.ElemType().Kind().String())
	buf.WriteString(">")
	first := true
	mLen := m.Len()
	for i := 0; i < mLen; i++ {
		if first {
			first = false
		} else {
			buf.WriteString(" ")
		}
		key := m.keys.Get(i)
		buf.WriteString(fmt.Sprintf("%v", key))
		buf.WriteString(":")
		buf.WriteString(fmt.Sprintf("%v", m.m[key]))
	}
	buf.WriteString("}")
	return buf.String()
}

type OrderedMap interface {
	Get(key interface{}) interface{}
	Put(key interface{}, elem interface{}) (interface{}, bool)
	Remove(key interface{}) interface{}
	Clear()
	Len() int
	Contains(key interface{}) bool
	FirstKey() interface{}
	LastKey() interface{}
	HeadMap(toKey interface{}) OrderedMap
	SubMap(fromKey interface{}, toKey interface{}) OrderedMap
	TailMap(fromKey interface{}) OrderedMap
	Keys() []interface{}
	Elems() []interface{}
	ToMap() map[interface{}]interface{}
	KeyType() reflect.Type
	ElemType() reflect.Type
}
type CompareFunc func(interface{}, interface{}) int8
type Keys interface {
	sort.Interface
	Add(k interface{}) bool
	Remove(k interface{}) bool
	Clear()
	Get(index int) interface{}
	GetAll() []interface{}
	Search(k interface{}) (index int, contains bool)
	ElemType() reflect.Type
	CompareFunc() CompareFunc
}

type myKeys struct {
	container   []interface{}
	compareFunc CompareFunc
	elemType    reflect.Type
}

func NewKeys(compareFunc CompareFunc, elemType reflect.Type) *myKeys {
	return &myKeys{container: make([]interface{}, 0), compareFunc: compareFunc, elemType: elemType}
}

func (m *myKeys) Len() int {
	return len(m.container)
}
func (m *myKeys) Less(i, j int) bool {
	return m.compareFunc(m.container[i], m.container[j]) == -1
}
func (m *myKeys) Swap(i, j int) {
	m.container[i], m.container[j] = m.container[j], m.container[i]
}
func (m *myKeys) Add(k interface{}) bool {
	if !m.isAcceptableElem(k) {
		return false
	}
	m.container = append(m.container, k)
	sort.Sort(m)
	return true
}
func (m *myKeys) Search(k interface{}) (index int, contains bool) {
	ok := m.isAcceptableElem(k)
	if !ok {
		return
	}
	index = sort.Search(m.Len(), func(i int) bool {
		return m.compareFunc(m.container[i], k) >= 0
	})
	if index < m.Len() && m.container[index] == k {
		contains = true
	}
	return
}
func (m *myKeys) Remove(k interface{}) bool {
	i, contains := m.Search(k)
	if contains == false {
		return false
	}
	m.container = append(m.container[:i], m.container[i+1:]...)
	return true
}
func (m *myKeys) isAcceptableElem(k interface{}) bool {
	if k == nil {
		return false
	}
	fmt.Println(k)

	if reflect.TypeOf(k) != m.elemType {
		return false
	}
	return true
}

func (m *myKeys) Clear() {
	m.container = make([]interface{}, 0)
}

func (m *myKeys) Get(index int) interface{} {
	if index >= m.Len() {
		return nil
	}
	return m.container[index]
}
func (m *myKeys) GetAll() []interface{} {
	initialLen := len(m.container)
	snapshot := make([]interface{}, initialLen)
	actualLen := 0
	for _, v := range m.container {
		if actualLen < initialLen {
			snapshot[actualLen] = v
		} else {
			snapshot = append(snapshot, v)
		}
		actualLen++
	}
	if actualLen < initialLen {
		snapshot = snapshot[:actualLen]
	}
	return snapshot
}

func (m *myKeys) ElemType() reflect.Type {
	return m.elemType
}
func (m *myKeys) CompareFunc() CompareFunc {
	return m.compareFunc
}
func (m *myKeys) String() string {
	var buf bytes.Buffer
	buf.WriteString("orderedMap {")
	buf.WriteString("orderedMap }")
	isFirst := true
	for _, v := range m.container {
		if isFirst {
			isFirst = false
		} else {
			buf.WriteString(" ")
		}
		buf.WriteString(fmt.Sprintf("%v", v))
	}
	buf.WriteString(" }")
	return buf.String()
}
