package redisdb

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"testing"
	"time"
)

func TestPushChannel(t *testing.T) {
}

type job struct {
	Name string
}

func TestDelayQueue(t *testing.T) {
	dq := newDelayQueue(Client, "lzh-delay-queue", func(msg *msg) bool {
		fmt.Println(msg.Data)
		j := job{}
		json.Unmarshal(msg.Data, &j)
		fmt.Println("Name",j.Name)
		return false
	})
	go dq.start()
	j := job{Name: "lzh"}
	data, _ := json.Marshal(j)
	dq.delay(&msg{
		Uuid: uuid.New().String(),
		Data: data,
		Try:  0,
	})
	time.Sleep(time.Hour)
}
