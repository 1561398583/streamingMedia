package gee

import (
	"fmt"
	"testing"
)

type U struct {
	name string
}

func TestMap(t *testing.T) {
	handlers :=make(map[string]*U)
	handlers["a"] = &U{name: "yxa"}
	h := handlers["b"]
	fmt.Println(h)
}

