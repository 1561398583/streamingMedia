package weberrors

import (
	"fmt"
	"testing"
)

func TestWebError(t *testing.T) {
	err := warpError()
	if err != nil {
		fmt.Println(err)
	}
}

func warpError()  error{
	err := getError()
	if err != nil {
		return Wrap(err)
	}
	return nil
}

func getError() error {
	return New("get a error", 1)
}

