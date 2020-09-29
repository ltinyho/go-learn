package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gogs.sharkgulf.cn/sg/library/qapp"
	"gogs.sharkgulf.cn/sg/library/qlog"
)

var log = qlog.WithField("app", "api")

func server(ctx context.Context) error {
	r := gin.Default()
	route(r)
	addr := viper.GetString("app.addr")
	return r.Run(addr)
}
func main() {
	qapp.New("api", func(a *qapp.Application) {
	}).AddDaemons(server).Run()
}
