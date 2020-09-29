package main

import (
	"encoding/base64"
	"fmt"
	"testing"
	"time"
)

type codeInfo struct {
	Val int
	Ts  time.Time
}

func TestSafeCode(t *testing.T) {
	valMap := make(map[int]codeInfo)
	i := 0
	count := 0
	for range time.Tick(time.Millisecond) {
		now := time.Now()
		code := safeCode("C0001", 0x01)
		if code > 500000 {
			fmt.Printf("code:%d\n", code)
		}
		oldCode, ok := valMap[code]
		if ok {
			i++
			milliseconds := now.Sub(oldCode.Ts).Milliseconds()
			if milliseconds > 1 {
				//fmt.Printf("code:%d,过期:%d\n old:%d now:%d", code, milliseconds, oldCode.Ts.UnixNano(), now.UnixNano())
				delete(valMap, code)
			} else {
				fmt.Printf("code:%d,重复:%d 重复之前:%d \n", code, i, count)
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

func TestName(t *testing.T) {
	squery := "a=afda=12321&asf=adsf&adf=adsfsd&adsf=adsfasd&asdf=basfsd"
	data := XOR([]byte(squery), []byte("921064e47128faf3febc6ce3836715ec"))
	dataStr := base64.StdEncoding.EncodeToString(data)
	fmt.Println(dataStr)
	baseDeData,_:=base64.StdEncoding.DecodeString(dataStr)
	fmt.Println()
	text := XOR(baseDeData, []byte("921064e47128faf3febc6ce3836715ec"))
	fmt.Println(string(text))
}
func XOR(input []byte, key []byte) []byte { //解密時僅需將原本的output改到input,key不變即可
	output := make([]byte, len(input))
	for i := range input {
		output[i] = input[i] ^ key[i%len(key)] //當input比key長時會不斷使用key對每一個byte加密
	}
	return output
}



