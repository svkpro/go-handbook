package main

import "fmt"

type MapPoint struct {
	Longitude, Latitude float64
}

func main() {
	var mp = map[string]MapPoint{
		"Bar": {
			53.906375, 27.544592,
		},
	}

	fmt.Println(mp["Bar"])
}

