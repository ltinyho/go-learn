package main

import (
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
