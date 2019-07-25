package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	fmt.Print(s)
	for _, v := range s {
		sum += v
	}
	c <- sum // from sum to channel
}

func main() {
	sequence := []int{7, 2, 8, -9, 4, 0}
	firstPart := sequence[:len(sequence)/2]
	secondPart := sequence[len(sequence)/2:]
	channel := make(chan int)
	go sum(firstPart, channel)
	go sum(secondPart, channel)
	x, y := <-channel, <-channel // from channel to x, y

	fmt.Println(x, y, x+y)
}

