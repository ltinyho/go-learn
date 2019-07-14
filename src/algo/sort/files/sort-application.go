package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

func main() {
	sort()
}

var (
	gPhonePrefixList []string = []string{"130", "131", "132", "133", "134", "135", "136", "137", "138",
		"139", "147", "150", "151", "152", "153", "155", "156", "157", "158", "159", "186",
		"187", "188"}
)

func sort() {
	data, err := ioutil.ReadFile("./numbers.txt")
	if err != nil {
		panic(err)
	}
	list := strings.Split(string(data), " ")
	length := 11
	for i := 0; i < length; i++ {
		bucketSort(list, length-i-1)
	}
	ioutil.WriteFile("./numbers-sort.txt", []byte(strings.Join(list, " ")), 0644)
}
func bucketSort(phones []string, idx int) {
	buckets := make([][]string, 10)
	for _, value := range phones {
		curStr := value[idx : idx+1]
		idx, _ := strconv.Atoi(curStr)
		buckets[idx] = append(buckets[idx], value)
	}
	i := 0
	for _, bucket := range buckets {
		for _, v := range bucket {
			phones[i] = v
			i++
		}
	}
}
func gen() {
	rand.Seed(time.Now().UnixNano())
	f, err := os.OpenFile("./numbers.txt", os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		panic(err)
	}
	for i := 0; i < 1000; i++ {
		n, err := f.WriteString(randPhones(100))
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(n)
		}
	}
}

const (
	gcNumberString = "1234567890"
)

func randPhones(length int) string {
	i := 0
	s := ""
	for i < length {
		s += randPhone() + " "
		i++
	}
	return s
}

func randPhone() string {
	s := ""
	l := 8
	for len(s) < l {
		s += rangeNums()
	}
	head := gPhonePrefixList
	return StringArray(&head) + s
}
func rangeNums() string {
	str := gcNumberString
	return RangeString(&str)
}

func StringArray(list *[]string) string {
	return (*list)[rand.Intn(len(*list))]
}

var m sync.Mutex

func RangeString(s *string) string {
	str := []rune(*s)
	index := rand.Intn(len(str))
	return string(str[index])
}
