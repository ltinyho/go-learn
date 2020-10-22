package zookeeper

import (
	"github.com/go-zookeeper/zk"
	"time"
)
var zkClient *zk.Conn
func init() {
	zkAddr := []string{"127.0.0.1:2181","127.0.0.1:2182","127.0.0.1:2183"}
	c, _, err := zk.Connect(zkAddr, time.Second) //*10)
	if err != nil {
		panic(err)
	}
	zkClient = c
}
