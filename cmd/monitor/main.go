package main

import (
	"bytes"
	"fmt"
	"gogs.sharkgulf.cn/sg/library/qlog"
	"io"
	"net"
	"time"
)

var log = qlog.WithField("app", "monitor")

func main() {
	go server()
	select {}
}
func server() {
	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Errorf("listen:%s", err)
		return
	}
	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Errorf("accept:%s", err)
		}
		go handleConn(conn)
	}
}
func handleConn(conn net.Conn) {
	defer conn.Close()
	addr := conn.RemoteAddr()
	result := bytes.NewBuffer(nil)
	var buf [10]byte
	for {
		n, err := conn.Read(buf[0:])
		fmt.Println("n", n)
		result.Write(buf[0:n])
		if err != nil {
			if err == io.EOF {
				continue
			} else {
				fmt.Println("read err:", err)
				break
			}
		} else {
			fmt.Println("recv:", addr, ":", result.String())
		}
		result.Reset()
	}
}

func client() {
	conn, err := net.Dial("tcp", ":1234")
	if err != nil {
		log.Error(err)
		return
	}
	for range time.Tick(time.Second) {
		_, err = conn.Write([]byte("hello"))
		if err != nil {
			log.Error(err)
			continue
		}
	}
}
