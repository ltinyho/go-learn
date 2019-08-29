package ctx

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestCtxTimeDone(t *testing.T) {
	tick := time.NewTicker(time.Second * 1)
	defer tick.Stop()
	myCtx := context.Background()
	var cancel context.CancelFunc
	myCtx, cancel = context.WithTimeout(myCtx, time.Second*3)
	defer cancel()
	go func() {
		time.Sleep(time.Second * 2)
		cancel()
	}()
	for {
		fmt.Println(time.Now())
		select {
		case val := <-tick.C:
			fmt.Println(val)
		case <-myCtx.Done():
			if myCtx.Err() == context.Canceled {
				fmt.Println("cancel")
			}
			return

		}
	}

}
func TestCtxTimePoll(t *testing.T) {
	myCtx := context.Background()
	var cancel context.CancelFunc
	myCtx, cancel = context.WithCancel(myCtx)
	defer cancel()
	for {
		fmt.Println(time.Now())
		select {
		case <-time.After(time.Second):
			fmt.Println(time.Now())
		case <-myCtx.Done():
			return
		}
	}
}

func TestConvertTime(t *testing.T) {
}
