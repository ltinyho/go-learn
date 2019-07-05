package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/ltinyho/go-learn/utils/redis"
	"log"
	"os"
	"time"
)

var RunMode string
var IsLocal bool

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	RunMode = os.Getenv("RUN_MODE")
	IsLocal = RunMode == "dev"
}
func tryLock(key string, resource string, duration time.Duration) (l *Lock, ok bool, err error) {
	l = &Lock{
		token:    uuid.New().String(),
		timeout:  duration,
		resource: resource,
	}
	ok, err = l.lock(key)
	return
}
func getExpressInfo(i int) {
	l, ok, err := tryLock(fmt.Sprintf("express_%v", i), "ltinyho", 10*time.Second)
	if err != nil {
		fmt.Println(err)
		return
	}
	if ok {
		time.Sleep(time.Second)
		fmt.Println("ok", i)
		l.unlock()
	} else {
		fmt.Println("lock", i)
	}
}

type Lock struct {
	token    string
	resource string
	timeout  time.Duration
}

func (l *Lock) lock(key string, ) (res bool, err error) {
	return redis.Client.SetNX(l.key(), l.token, l.timeout).Result()
}
func (l *Lock) key() string {
	return fmt.Sprintf("redislock:%s", l.resource)
}
func (l *Lock) unlock() {
	redis.Client.Del(l.key())
}
func (l *Lock) exec(fn func() error) {
	ttl, err := redis.Client.TTL(l.key()).Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ttl)
}

func main() {
	fmt.Println(uuid.New())
	for i := 0; i < 1; i++ {
		for j := 0; j < 5; j++ {
			go getExpressInfo(i)
		}
	}
	time.Sleep(time.Hour)
}
