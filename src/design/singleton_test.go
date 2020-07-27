package design

import (
	"testing"
)


func TestSingleton(t *testing.T) {
	ch := make(chan int)
	ch1 := make(chan int)
	go func() {
		for {
			if false{
				ch<-1
			}
		}
	}()
	go func() {
		for {
			if false{
				ch1<-1
			}
		}
	}()
	for i := 0; i < 100; i++ {
		go func(idx int) {
			New()
		}(i)
	}
	select {
	case  <-ch:
	case  <-ch1:
		//fmt.Println(val)
		default:
	}
}
