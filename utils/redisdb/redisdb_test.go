package redisdb

import (
	"fmt"
	"github.com/go-redis/redis"
	"math/rand"
	"testing"
	"time"
)

func TestRedisSet(t *testing.T) {
	data := make([]byte, 1<<10)
	rand.Read(data)
	for range time.Tick(time.Millisecond) {
		i := rand.Int31n(1 << 30)
		key := fmt.Sprintf("lzh%d", i)
		Client.Set(key, data, -1)
	}
}

func TestSubscribe(t *testing.T) {
	go func() {
		pub := Client.Subscribe("test")
		for {
			_, err := pub.ReceiveMessage()
			if err != nil {
				t.Error(err)
			}
		}
	}()
	go func() {
		pub := Client.Subscribe("test1")
		for {
			_, err := pub.ReceiveMessage()
			if err != nil {
				t.Error(err)
			}
			time.Sleep(time.Second*5)
			Client.Publish("test",redis.Pong{Payload:"1" })
		}
	}()
	time.Sleep(time.Hour)
}
