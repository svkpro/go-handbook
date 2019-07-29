package main

import "fmt"

var s = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for i, v := range s {
		fmt.Printf("%d = %d\n", i, v)
	}
}