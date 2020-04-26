package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"testing"
	"time"
)

// goroutine 泄露
func TestLeak(t *testing.T) {

}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	go func() {
		for n := range workLeak(ctx) {
			if n == 2 {
				//cancel() // 通过 context cancel 通知 goroutine
				break // goroutine 泄露
			}
		}
	}()
	go func() {
		log.Println(http.ListenAndServe("localhost:8080", nil))
	}()
	time.Sleep(time.Hour)
}

func workLeak(ctx context.Context) <-chan int {
	ch := make(chan int)
	go func() {
		var n int
		for {
			select {
			case <-ctx.Done():
				fmt.Println("done")
				return
			default:
				ch <- n
				n++
				fmt.Println(n)
				time.Sleep(time.Millisecond * 50)
			}

		}
	}()
	return ch
}
