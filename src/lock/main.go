package lock

import (
	"github.com/ltinyho/go-learn/utils/redis"
	"time"
)

type Lock interface {
	Lock()
	UnLock()
}

type MyLock struct {
	isLock bool
}

func (l *MyLock) Lock() {
	l.isLock = true
	for l.isLock {
		time.Sleep(time.Nanosecond)
	}
}

func (l *MyLock) UnLock() {
	l.isLock = false
}

func main() {
	redis.Client.Set("lzh", "2", time.Second*10)

}
