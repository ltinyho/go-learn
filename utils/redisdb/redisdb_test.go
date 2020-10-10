package redisdb

import (
	"encoding/json"
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
			time.Sleep(time.Second * 5)
			Client.Publish("test", redis.Pong{Payload: "1"})
		}
	}()
	time.Sleep(time.Hour)
}

func TestLock(t *testing.T) {
	for i := 0; i < 1000; i++ {
		go getLock()
	}
	select {}
}

type lockRes struct {
	Name string
	Tag  int64
}

// 分布式锁
func getLock() {
	tag := time.Now().UnixNano()
	key := "lzh:lock"
	a := map[string]interface{}{
		"Name": "ok",
		"tag":  tag,
	}
	val, _ := json.Marshal(a)
	_, err := Client.SetNX(key, val, 5*time.Second).Result()
	if err != nil {
		fmt.Printf("lock:%s", err)
	}
	second := rand.Int31n(10)
	time.Sleep(time.Second * time.Duration(second))
	res, err := Client.Get(key).Result()
	if err != nil {
		fmt.Printf("res:%s", err)
	}
	lr := lockRes{}
	json.Unmarshal([]byte(res), &lr)
	if lr.Tag == tag {
		Client.Del(key)
	} else {
		fmt.Println("tag not equal")
	}
}

func TestName(t *testing.T) {
	for i := 0; i < 513; i++ {
		Client.HSet("lzh",fmt.Sprintf("%d",i),i)
	}
}
