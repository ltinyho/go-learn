package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"testing"
	"time"
)

func TestSelect(t *testing.T) {
	Client, err := gorm.Open(mysql.Open("root:ltinyho@tcp(localhost:3306)/ltinyho?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{
	})
	db, err := Client.DB()
	db.SetMaxOpenConns(10)
	if err != nil {
		panic(err)
	}
	go func() {
		log.Println(http.ListenAndServe(":6060", nil))
	}()
	for range time.Tick(time.Millisecond) {
		rows, err := Client.Raw("select version()").Rows()
		if err != nil {
			fmt.Println(err)
			continue
		}
		rows.Close()
	}
	time.Sleep(time.Hour)
}
type out struct {
	filename string
	f        *os.File
}

func newOut(filename string) *out {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}
	return &out{
		filename: filename,
		f:        f,
	}
}

func (o *out) Write(p []byte) (n int, err error) {
	return o.f.Write(p)
}

func TestName(t *testing.T) {
	log := logrus.NewEntry(func() *logrus.Logger {
		entry := logrus.New()
		entry.SetLevel(logrus.DebugLevel)
		entry.SetReportCaller(true)
		entry.SetFormatter(&logrus.JSONFormatter{
		})
		entry.SetOutput(newOut("/Users/ltinyho/logs/filebeat/test1.log"))
		return entry
	}())


	for range time.Tick(time.Second) {
		log.Info("ok")
	}
}
