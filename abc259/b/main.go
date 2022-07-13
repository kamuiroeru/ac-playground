package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, d float64
	fmt.Scanf("%f %f %f", &a, &b, &d)
	// fmt.Println(a, b, d)
	dRad := d * math.Pi / 180.0
	retA := a*math.Cos(dRad) - b*math.Sin(dRad)
	retB := a*math.Sin(dRad) + b*math.Cos(dRad)
	fmt.Println(retA, retB)
}
