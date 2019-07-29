package main

import "fmt"

func main() {
	a := make([]int, 5)
	printMadeSlice("a", a)

	b := make([]int, 0, 5)
	printMadeSlice("b", b)

	c := b[:2]
	printMadeSlice("c", c)

	d := c[2:5]
	printMadeSlice("d", d)
}

func printMadeSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

