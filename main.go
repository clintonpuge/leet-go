package main

import (
	"fmt"
)

type Point struct {
	X int
	Y int
}

func main() {
	fmt.Println("Codes from Leet Code exercises.")

	lines := []Point{
		{X: 1, Y: 1},
		{X: 2, Y: 2},
		{X: 3, Y: 3},
	}
	max := 0
	for _, point := range lines {
		if point.Y > max {
			max = point.Y
		}
	}
	fmt.Println(max)
}
