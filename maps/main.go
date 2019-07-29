package main

import "fmt"

type TimeZone struct {
	Zone, Offset string
}

func main() {
	tz := make(map[string]TimeZone)

	tz["BY"] = TimeZone{
		"Europe/Minsk", "UTC +03:00",
	}
	fmt.Println(tz["BY"])
}