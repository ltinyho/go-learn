package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"gogs.sharkgulf.cn/sg/library/utils/qutil"
	"net/url"
	"strconv"
)

func main() {

}

const xorNum = 0xFEA



func offlineSign(query string) (sign string) {
	return qutil.QMD5String([]byte("bs" + query + "offline"))
}

func parseQueryQR(cipher string, key string) (data url.Values, err error) {
	query, err := base64.StdEncoding.DecodeString(cipher)
	qr := xor(query, []byte(key))
	data, err = url.ParseQuery(string(qr))
	if err != nil {
		return nil, fmt.Errorf("parse query error,%w", err)
	}
	qSign := data.Get("sign")
	data.Del("sign")
	sign := offlineSign(data.Encode())
	if qSign != sign {
		err = fmt.Errorf("sign not equal %s!=%s", qSign, sign)
		return nil, err
	}
	return
}

func genQueryQR(queryStr string, key string) (qrCode string, err error) {
	data, err := url.ParseQuery(queryStr)
	if err != nil {
		return
	}
	data.Del("sign")
	encode := data.Encode()
	query := encode + "&sign" + offlineSign(encode)
	d := xor([]byte(query), []byte(key))
	return base64.StdEncoding.EncodeToString(d), nil
}

func createSafeCode(cid string, randNum, state int) (code string) {
	codeNum := genSafeCode(cid, randNum, state)
	return encryptCode(codeNum)
}

func getSafeCodeData(code string) (cidNum, randNum, state int) {
	return deSafeCode(decryptCode(code))
}

func encryptCode(code int) string {
	code ^= xorNum
	valStr := fmt.Sprintf("%06d", code) // 小于 6 位添加前导 0
	return exchangeNum(valStr)
}

func decryptCode(code string) int {
	code = exchangeNum(code)
	codeNum, _ := strconv.Atoi(code)
	return codeNum ^ xorNum
}

func genSafeCode(cid string, randNum, state int) (code int) {
	cidNum := getCidNum(cid)
	str := ""
	str += fmt.Sprintf("%0.11b", randNum)
	str += fmt.Sprintf("%0.4b", cidNum)
	str += fmt.Sprintf("%0.4b", state)
	val, _ := strconv.ParseInt(str, 2, 32)
	code = int(val)
	return
}

// 1.交换奇偶数字
// 2.转成 int与 0xFEA 异或
// 3.转成 2 进制,按位取值
func deSafeCode(code int) (cidNum, randNum, state int) {
	codeStr := fmt.Sprintf("%019b", code)
	randBits := codeStr[0:11]
	cidBits := codeStr[11:15]
	stateBits := codeStr[15:19]
	_randNum, _ := strconv.ParseInt(randBits, 2, 32)
	randNum = int(_randNum)
	_cidNum, _ := strconv.ParseInt(cidBits, 2, 32)
	cidNum = int(_cidNum)
	_state, _ := strconv.ParseInt(stateBits, 2, 32)
	state = int(_state)
	return
}

// 柜子 cid 字符串转成 ascii 码累加后对 2^4 取余
func getCidNum(cid string) (num int) {
	for _, _cid := range cid {
		num += int(_cid)
	}
	return num % (1 << 4)
}
func exchangeNum(valStr string) string {
	// 将数字奇数位和偶数位互换,比如: 49 42 86 => 94 24 68
	codeBytes := bytes.NewBufferString("")
	codeBytes.WriteByte(valStr[1])
	codeBytes.WriteByte(valStr[0])
	codeBytes.WriteByte(valStr[3])
	codeBytes.WriteByte(valStr[2])
	codeBytes.WriteByte(valStr[5])
	codeBytes.WriteByte(valStr[4])
	return codeBytes.String()
}

func xor(input []byte, key []byte) []byte { // 解密时仅需将原本的output改到input,key不变即可
	output := make([]byte, len(input))
	for i := range input {
		output[i] = input[i] ^ key[i%len(key)] //当input比key长时会不断使用key对每一个byte加密
	}
	return output
}
