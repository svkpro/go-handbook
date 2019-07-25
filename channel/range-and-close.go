package main

import (
	"fmt"
)

func fibonacci(n int, channel chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		channel <- x
		x, y = y, x+y
	}
	close(channel)
}

func main() {
	channel := make(chan int, 10)
	go fibonacci(cap(channel), channel)
	for i := range channel {
		fmt.Println(i)
	}
}
