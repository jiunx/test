package main

import "fmt"

func main() {
	arr1 := make([]int, 2)
	arr1[0] = 1
	arr1[1] = 2
	fmt.Println(arr1)
	add(arr1)
	fmt.Println(arr1)
}

func add(arr []int) {
	for _, i2 := range arr {
		i2++
	}
}
