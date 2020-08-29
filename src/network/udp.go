package main

import (
	"encoding/binary"
	"fmt"
	"github.com/sirupsen/logrus"
	"net"
	"os"
	"time"
)

var (
	log = logrus.NewEntry(func() *logrus.Logger {
		return &logrus.Logger{
			Out:          os.Stdout,
			Hooks:        nil,
			Formatter:    &logrus.JSONFormatter{},
			ReportCaller: true,
			Level:        logrus.DebugLevel,
			ExitFunc:     nil,
		}
	}())
)

func main() {
	ch := make(chan int, 0)
	go func() {
		addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8888")
		if err != nil {
			log.Error(err)
			return
		}
		conn, err := net.ListenUDP("udp", addr)
		if err != nil {
			log.Error(err)
			return
		}
		go func() {
			for {
				buf := make([]byte, 1024)
				n, add, err := conn.ReadFrom(buf)
				if err != nil {
					log.Error(err)
					continue
				}
				fmt.Printf("data:%x\n", buf[:n])
				u := binary.LittleEndian.Uint16(buf[:n])
				u1 := binary.BigEndian.Uint16(buf[:n])
				fmt.Printf("data:%x,%d\n", u,u)
				fmt.Printf("datau1:%x,%d\n", u1,u1)
				_, err = conn.WriteTo(buf[:n], add)
				if err != nil {
					log.Error(err)
					continue
				}
			}
		}()
		ch <- 1
	}()
	conn, err := net.Dial("udp", ":8888")
	if err != nil {
		log.Error(err)
		return
	}
	data := make([]byte, 4)
	binary.LittleEndian.PutUint16(data[0:2], 0xabcd)
	binary.LittleEndian.PutUint16(data[2:4], 0xef01)
	fmt.Println("%x",data)
	_, err = conn.Write(data)
	if err != nil {
		log.Error(err)
	}
	time.Sleep(time.Hour)
}
