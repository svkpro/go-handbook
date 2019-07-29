package main

import "fmt"

func main() {
	var s []int
	printTheSlice(s)

	// append works on nil slices.
	s = append(s, 0)
	printTheSlice(s)

	// The slice grows as needed.
	s = append(s, 1)
	printTheSlice(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printTheSlice(s)
}

func printTheSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}