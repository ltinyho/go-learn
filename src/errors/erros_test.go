package errors

import (
	"errors"
	"fmt"
	"os"
	"testing"
)

func TestError(t *testing.T) {
	fmt.Println()
	err := errors.New("haha")
	switch err := err.(type) {
	case CacheError:
		fmt.Println(err.Error())
	}
	os.IsExist()
}
