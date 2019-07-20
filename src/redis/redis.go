package main

import (
	"github.com/ltinyho/go-learn/utils/redisdb"
	"sync/atomic"
)
import "fmt"

func main() {
	mq := redisdb.NewConsumer(redisdb.ConsumerParams{
		Channel: redisdb.OrderChannel,
		DoFunc: func(m interface{}) {
			switch msg := m.(type) {
			case *redisdb.OrderMessage:
				go func(id int) {
					fmt.Println(id)
				}(msg.Id)
			}
		},
		Message: &redisdb.OrderMessage{},
		Mode:    redisdb.PushPop,
	})
	mq.Listen()
}

var start uint64 = 0
var count uint64 = 0

func setKeys(num int) {
	setKeysPipe(num)
	return
	pipe := redisdb.Client.Pipeline()
	for i := 0; i < num; i++ {
		key := fmt.Sprintf("lzh-%v", i)
		pipe.HSet(key, fmt.Sprintf("%d", i), -1)
	}
	pipe.Exec()
}

func setKeysPipe(num int) {
	pipe := redisdb.Client.Pipeline()
	for num > 0 {
		atomic.AddUint64(&count, 1)
		key := fmt.Sprintf("lzh-%v", count)
		pipe.Set(key, fmt.Sprintf("%d", count), -1)
		num--
	}
	pipe.Exec()
}
