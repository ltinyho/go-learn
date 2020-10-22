package main

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"github.com/go-zookeeper/zk"
	"gogs.sharkgulf.cn/sg/library/qapp"
	"gogs.sharkgulf.cn/sg/library/qlog"
)

var (
	log      = qlog.WithField("app", "api")
	zkClient *zk.Conn
)

func server(ctx context.Context) error {
	r := gin.Default()
	route(r)
	addr := viper.GetString("app.addr")
	return r.Run(addr)
}

func zkInitFunc(ctx context.Context) (qapp.CleanFunc, error) {
	zkAddr := viper.GetStringSlice("zk.addr")
	log.Debugf("addr:%v", zkAddr)
	c, _, err := zk.Connect(zkAddr, time.Second) //*10)
	if err != nil {
		return nil, err
	}
	zkClient = c
	return func(ctx context.Context) {
		c.Close()
	}, nil
}
func main() {
	qapp.New("api", func(a *qapp.Application) {
	}).AddInitStage("zkInit", zkInitFunc).AddDaemons(server).Run()
}
