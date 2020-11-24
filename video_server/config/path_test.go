package config

import (
	"fmt"
	"os"
	"testing"
)

func TestPath(t *testing.T)  {
	dirString, err := os.Getwd()
	if err != nil{
		t.Error(err)
	}
	fmt.Println(dirString)
}
