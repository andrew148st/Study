package main

import (
	"fmt"
	"math"
)

type point struct {
	x int
	y int
}

func (p point) distanceTo(p2 point) int {
	return int(
		math.Sqrt(
			math.Pow(float64(p2.x-p.x), 2) + math.Pow(float64(p2.y-p.y), 2),
		),
	)
}
func main() {
	var deliverTo point
	deliverTo.x = 20
	deliverTo.y = 30
	fmt.Printf("Delivering to: x= %d y= %d\n", deliverTo.x, deliverTo.y)
	deliverFrom := point{x: 70, y: 50}
	fmt.Printf("Delivering from: x= %d y= %d\n", deliverFrom.x, deliverFrom.y)
	distance := deliverFrom.distanceTo(deliverTo)
	fmt.Printf("Distance: %d\n", distance)
}
