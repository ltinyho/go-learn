package test_method

import (
	"fmt"
	"reflect"
	"testing"
)

type Cache interface {
	Get(interface{}) interface{}
}

type mock struct{}

func (mock) Get(interface{}) interface{} {
	panic("implement me")
}

func createCacheKey(inter interface{}) string {
	v := reflect.ValueOf(inter)
	t := reflect.TypeOf(inter)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	} else {
		return ""
	}
	return t.Name()
}

func TestGetCache(t *testing.T) {
	key := createCacheKey(&mock{})
	fmt.Println(key)
}
