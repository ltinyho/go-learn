package _json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDecode(t *testing.T) {
	type Message struct {
		Id json.Number `json:"num_integer"`
	}
	var message = `{
		"num_integer": 10,
		"num_float": 10.5
	}`
	var msg Message
	if err := json.Unmarshal([]byte(message), &msg); err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", msg)

	numInt, err := msg.Id.Int64()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%T",numInt)
}
