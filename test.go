package main

import (
	"fmt"
	"strings"
)

func main() {
	var a strings.Builder
	for i := 0; i < 55; i++ {
		a.WriteByte('a')
	}
	fmt.Println(a.String())
}
