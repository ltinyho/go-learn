package main

import (
	"github.com/ltinyho/go-learn/utils/redisdb"
	"testing"
	"time"
)

func TestSetKeys(t *testing.T) {
	var i = 0
	for {
		redisdb.SendOrder(redisdb.OrderMessage{
			Id: i,
		})
		i++
		time.Sleep(time.Microsecond)
	}
}

func BenchmarkSetKeys(b *testing.B) {
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		setKeys(1000)
	}
	b.StopTimer()
}
