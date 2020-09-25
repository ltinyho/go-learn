package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {

}

func safeCode(cid string, res int) (code int) {
	cidNum := 0
	for _, _cid := range cid {
		cidNum += int(_cid)
	}
	cidMod := cidNum % 1 << 4
	rand.Seed(time.Now().UnixNano())
	randNum := rand.Intn(1 << 10)
	valStr:=""
	valStr+=fmt.Sprintf("%0.11b", randNum)
	valStr+=fmt.Sprintf("%0.4b", cidMod)
	valStr+=fmt.Sprintf("%0.4b", res)
	val,err:=strconv.ParseInt(valStr,2,32)
	if err != nil {
		return 0
	}
	return int(val)
}
