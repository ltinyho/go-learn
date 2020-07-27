package design

import (
	"fmt"
	"sync"
)

type singleton struct {
}

var single *singleton

var lock sync.Mutex

func New() *singleton {
	if single == nil {
		fmt.Println("enter singleton")
		lock.Lock()
		defer lock.Unlock()
		if single == nil {
			fmt.Println("new singleton")
			single = &singleton{}
		}
	} else {
	}
	return single
}
