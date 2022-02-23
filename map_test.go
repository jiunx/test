package main

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	var m = make(map[string]int)
	m["a"] = 1
	m["b"] = 2
	if _, ok := m["b"];ok {
		fmt.Println("ok")
	}
}
