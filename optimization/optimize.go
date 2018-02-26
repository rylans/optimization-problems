// One-dimensional function optimization
package main

import (
	"fmt"
	"github.com/rylans/chromosomes/optimize"
	"math"
)

func optimize1d() {
	max, min := 3, -3
	fmt.Println("Bounded function optimization in range", min, "to", max)
	fmt.Println("f(x) = (x^2 + x) * cos(x)")

	optimizef := func(x float64) float64 {
		return ((x * x) + x) * math.Cos(x)
	}

	result := optimize.BoundedMaximize(optimizef, min, max)

	fmt.Println("Best result: X=", result)
	fmt.Println("f(x)=", optimizef(result))
}

func optimize2d() {
	max, min := 2, -2
	fmt.Println("Bounded function optimization in range", min, "to", max)
	fmt.Println("f(x,y) = cos(x)cos(y)e^(-x^2 -y^2) ")

	optimizef := func(x, y float64) float64 {
		return math.Cos(x) * math.Cos(y) * math.Exp(-(x*x)-(y*y))
	}

	resX, resY := optimize.BoundedMaximize2D(optimizef, min, max)

	fmt.Println("Best result: (X,Y)=", resX, resY)
	fmt.Println("f(x,y)=", optimizef(resX, resY))
}

func main() {
	optimize1d()
	fmt.Println()
	optimize2d()
}
