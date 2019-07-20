package redisdb

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"log"
)

type ConsumerMode string

const (
	PushPop ConsumerMode = "PushPop"
	PubSub  ConsumerMode = "PubSub"
)

type Channel string

const (
	OrderChannel Channel = "ltinyho.top.order"
)

func PushChannel(channel Channel, val interface{}) {
	result, err := json.Marshal(val)
	if err != nil {
		fmt.Printf("Push channel json err: %v", err)
		return
	}
	res, err := Client.RPush(string(channel), string(result)).Result()
	if err != nil {
		fmt.Printf("Push channel err: %v", err)
		return
	} else {
		fmt.Println(res)
	}
}

type OrderMessage struct {
	Id int `json:"id"`
}

func SendOrder(val OrderMessage) {
	PushChannel(OrderChannel, val)
}

type Consumer struct {
	Mode    ConsumerMode
	Message interface{}
	DoFunc  func(m interface{})
	Channel Channel
	c       *redis.Client
}
type ConsumerParams struct {
	Mode    ConsumerMode
	Message interface{}
	DoFunc  func(m interface{})
	Channel Channel
}

func NewConsumer(params ConsumerParams) *Consumer {
	con := &Consumer{
		Mode:    params.Mode,
		Message: params.Message,
		DoFunc:  params.DoFunc,
		Channel: params.Channel,
		c:       Client,
	}
	con.check()
	return con
}

// 检测实例参数
func (c *Consumer) check() bool {
	if c.Channel == "" {
		log.Fatal("channel empty")
		return false
	}
	if c.Mode != PushPop && c.Mode != PubSub {
		log.Fatal("mod error")
		return false
	}
	return true
}
func (c *Consumer) pop() () {
	for {
		result, err := c.c.LPop(string(c.Channel)).Bytes()
		if err != nil {
			fmt.Printf("pop %v", err)
			continue
		}
		err = json.Unmarshal(result, &c.Message)
		if err != nil {
			fmt.Printf("json error %v", err)
			continue
		}
		c.do()
	}
}
func (c *Consumer) sub() () {
	sub := c.c.Subscribe(string(c.Channel))
	for {
		result, err := sub.ReceiveMessage()
		if err != nil {
			fmt.Printf("json error %v", err)
			continue
		}
		err = json.Unmarshal([]byte(result.Payload), &c.Message)
		if err != nil {
			fmt.Printf("json error %v", err)
			continue
		}
		c.do()
	}
}

func (c *Consumer) do() () {
	c.DoFunc(c.Message)
}
func (c *Consumer) Listen() {
	fmt.Printf("mq listen %v", c.Channel)
	if c.Mode == PushPop {
		c.pop()
	} else if c.Mode == PubSub {
		for {
			c.sub()
		}
	}
}
