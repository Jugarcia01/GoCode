package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func main() {

	var getDistance = func(p1, p2 Point) float64 {
		var d1 = math.Pow(p1.x-p2.x, 2)
		var d2 = math.Pow(p1.y-p2.y, 2)
		return math.Sqrt(d1 + d2)
	}

	var point1 = Point{0.0, 1.0}
	var point2 = Point{5.5, 10.22}

	var distance = getDistance(point1, point2)
	fmt.Printf("La distancia entre los puntos %v y %v es: %.2f\n", point1, point2, distance)

}
