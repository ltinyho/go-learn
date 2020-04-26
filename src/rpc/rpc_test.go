package main

import (
	"fmt"
	"log"
	"testing"
)

func TestRpcClient(t *testing.T) {
	client,err:=DialHelloService("tcp","localhost:1233")
	if err != nil {
		log.Fatal("dialing:",err)
	}
	var reply string
	err= client.Hello("lzh",&reply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply)
}
