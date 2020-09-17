package main

import (
	"fmt"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

func main(){
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


