package redisdb

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/google/uuid"
	"log"
	"strconv"
	"time"
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
func (c *Consumer) pop() {
	for {
		result, err := c.c.BLPop(0, string(c.Channel)).Result()
		if err != redis.Nil && err != nil {
			fmt.Printf("pop %v", err)
			continue
		}
		if len(result) <= 0 {
			continue
		}
		err = json.Unmarshal([]byte(result[1]), &c.Message)
		if err != nil {
			fmt.Printf("json error %v", err)
			continue
		}
		c.do()
	}
}
func (c *Consumer) sub() {
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

func (c *Consumer) do() {
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

type delayQueue struct {
	client *redis.Client
	key    string
	fn     func(msg *msg) bool
}
type msg struct {
	Uuid string
	Data []byte
	Try  int
}

func newDelayQueue(client *redis.Client, key string, fn func(msg *msg) bool) *delayQueue {
	return &delayQueue{client: client, key: key, fn: fn}
}

func (d *delayQueue) start() {
	for {
		ts := time.Now().Unix()
		tsStr := strconv.Itoa(int(ts))
		data, err := d.client.ZRangeByScore(d.key, redis.ZRangeBy{
			Max:    tsStr,
			Min: "-inf",
			Offset: 0,
			Count:  1,
		}).Result()
		if err == redis.Nil || len(data) == 0 {
			fmt.Print("empty\n")
			time.Sleep(time.Second)
			continue
		}
		// 通过 Rem 保证多个 goroutine 不会同时处理
		// TODO 使用 lua 脚本,使用服务端原子操作,解决多客户端抢占问题.
		ok, err := d.client.ZRem(d.key, data[0]).Result()
		if err != nil {
			fmt.Printf("err:%s", err)
			continue
		}
		if ok == 1 {
			m := &msg{}
			err = json.Unmarshal([]byte(data[0]), m)
			if err != nil {
				fmt.Printf("err:%s", err)
				continue
			}
			success := d.fn(m)
			if success == false && m.Try <= 3 {
				d.delay(m)
			}
		}
	}
}

func (d *delayQueue) delay(msg *msg) {
	msg.Try += 1
	msg.Uuid = uuid.New().String()
	data, err := json.Marshal(*msg)
	if err != nil {
		fmt.Println("err",err)
		return
	}
	ok, err := d.client.ZAdd(d.key, redis.Z{
		Score:  float64(time.Now().Add(time.Second * 5).Unix()),
		Member: string(data),
	}).Result()
	if err != nil {
		fmt.Println("err", err)
		return
	}
	fmt.Printf("ok:%d\n", ok)
}
