package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"math/rand"
	"net/url"
	"os"
	"strconv"
	"testing"
	"time"
)

type codeInfo struct {
	Val string
	Ts  time.Time
}

func TestCreateSafeCode(t *testing.T) {
	valMap := make(map[string]codeInfo)
	i := 0
	count := 0
	for range time.Tick(time.Millisecond) {
		rand.Seed(time.Now().UnixNano())
		now := time.Now()
		code := createSafeCode("C0001", rand.Intn(1<<11), 0x01)
		oldCode, ok := valMap[code]
		fmt.Println(code)
		if ok {
			i++
			milliseconds := now.Sub(oldCode.Ts).Milliseconds()
			if milliseconds > 1 {
				//fmt.Printf("code:%d,过期:%d\n old:%d now:%d", code, milliseconds, oldCode.Ts.UnixNano(), now.UnixNano())
				delete(valMap, code)
			} else {
				fmt.Printf("code:%s,重复:%d 重复之前:%d \n", code, i, count)
				count = 0
			}
		} else {
			count++
			_code := codeInfo{
				Val: code,
				Ts:  time.Now(),
			}
			valMap[code] = _code
		}
	}
}

func TestXorName(t *testing.T) {
	query := "a=afda=12321&asf=adsf&adf=adsfsd&adsf=adsfasd&asdf=basfsd"
	data := xor([]byte(query), []byte("921064e47128faf3febc6ce3836715ec"))
	dataStr := base64.StdEncoding.EncodeToString(data)
	fmt.Println(dataStr)
	baseDeData, _ := base64.StdEncoding.DecodeString(dataStr)
	text := xor(baseDeData, []byte("921064e47128faf3febc6ce3836715ec"))
	fmt.Println(string(text))
}
func TestSafeCode(t *testing.T) {
	cid := "C0001"
	rn := 44
	s := 0xf
	code := genSafeCode(cid, rn, s)
	cidNum, randNum, state := deSafeCode(code)
	assert.Equal(t, cidNum, getCidNum(cid))
	assert.Equal(t, randNum, rn)
	assert.Equal(t, state, s)
}

func TestEncrypt(t *testing.T) {
	cid := "C0001"
	rn := 44
	s := 0xf
	code := createSafeCode(cid, rn, s)
	cidNum, randNum, state := deSafeCode(decryptCode(code))
	assert.Equal(t, cidNum, getCidNum(cid))
	assert.Equal(t, randNum, rn)
	assert.Equal(t, state, s)
}
func TestDe(t *testing.T) {
	max := 1 << 11
	for i := 0; i < max; i++ {
		code := createSafeCode("C0002", i, 0x00)
		fmt.Println(code)
	}
	fmt.Printf("%019b\n", xorNum)
}

func TestPz(t *testing.T) {
	cid := 4
	rn := 1687
	s := 1
	f, err := os.Open("./safecode.txt")
	if err != nil {
		t.Error(err)
		return
	}
	buf := bufio.NewReader(f)
	for {
		line, _, err := buf.ReadLine()
		if err == io.EOF {
			break
		}
		i := 1
		for {
			lineNum, _ := strconv.Atoi(string(line))
			code := lineNum + i
			cidNum, randNum, state := deSafeCode(decryptCode(fmt.Sprintf("%06d", code))) // 614115 018272 342299
			if cidNum == cid && randNum == rn && state == s {
				fmt.Printf("i:%d,code:%d\n", i, code)
				break
			}
			i++
		}
	}
}

func TestXorNum(t *testing.T) {
	//fmt.Println(0b000 0000 0000 0100 0000 ^ xorNum)
	fmt.Println(0x00040 ^ xorNum)
	fmt.Println(0x00140 ^ xorNum)
	fmt.Printf("%019b\n", 0x00040)
	fmt.Printf("%019b\n", 0xFEA)
	fmt.Printf("%019b\n", 4010)
	fmt.Println()
	fmt.Printf("%019b\n", 0x00140)
	fmt.Printf("%019b\n", 0xFEA)
	fmt.Printf("%019b\n", 3754)
}

func TestParseQuery(t *testing.T) {
	data, err := parseQueryQR("WltVDXUEVQQGF1tWOQMHRxI6AQJGAgZaTEoLAgEFQwpXbVNRQkBYdlZFRghXRw9dORULBwtTVBVIXEZoU1QRF2ZRUEBXVwxATgwKCFZHFlwWOgAKUl4nUkxHBgUXRQoTZkJYVAsEVxJFUFxcOQ8TXltWWloAU1ACDAUBABdHAAJKXV8NBxIRRwoABAhUU1QLU11XRUUKAl0FBA8FU1FSUQ4GAFQEBVMFUQNRXVRTUwQFUwBWVFEHClo=", "921064e47128faf3febc6ce3836715ec")
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(data.Encode())
}

func TestGenQuery(t *testing.T) {
	query := url.Values{}
	query.Add("cid", "C0001")
	query.Add("reason", "1")
	query.Add("in_bid", "Batt01")
	query.Add("in_pid", "1")
	query.Add("in_batt_capacity", "500")
	query.Add("pop_bid", "Batt02")
	query.Add("pop_pid", "2")
	query.Add("pop_batt_capacity", "800")
	query.Add("rand_num", "1234")
	query.Add("ts", "1602228585")
	query.Add("ver", "1")
	sign:=offlineSign(query.Encode())
	fmt.Println(query.Encode()+"&sign"+sign)
}

//cid=C0001&in_batt=Batt01&in_batt_capacity=500&in_pid=01&pop_batt_capacity=800&pop_bid=Batt02&pop_pid=02&rand_num=38960514677&reason=1&ts=1602228585

//cid=C0001&in_batt_capacity=500&in_batt=Batt01&in_pid=01&pop_batt_capacity=800&pop_bid=Batt02&pop_pid=02&rand_num=38960514677&reason=1&ts=1602228585
