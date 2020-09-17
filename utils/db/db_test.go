package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	_ "net/http/pprof"
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
