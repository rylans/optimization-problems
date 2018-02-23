// One-dimensional function optimization
package main

import (
	"fmt"
	"github.com/rylans/chromosomes"
	"github.com/rylans/chromosomes/optimize"
	"math"
)

func rescale(val uint8, min int, max int) float64 {
	newRange := float64(max) - float64(min)
	return (newRange/255.0)*(float64(val)-255.0) + float64(max)
}

type optimizeFn1d func(x float64) float64
type optimizeFn2d func(x, y float64) float64

func boundedFitnessFunc1d(min int, max int, optimizef optimizeFn1d) func(c *chromosomes.Chromosome) float64 {
	f := func(c *chromosomes.Chromosome) float64 {
		realVal := rescale(c.Get("X"), min, max)
		return optimizef(realVal)
	}
	return f
}

func boundedFitnessFunc2d(min int, max int, optimizef optimizeFn2d) func(c *chromosomes.Chromosome) float64 {
	f := func(c *chromosomes.Chromosome) float64 {
		realValX := rescale(c.Get("X"), min, max)
		realValY := rescale(c.Get("Y"), min, max)
		return optimizef(realValX, realValY)
	}
	return f
}

func optimize1d() {
	max, min := 10, -10
	fmt.Println("Bounded function optimization in range", min, "to", max)
	fmt.Println("f(x) = (x^2 + x) * cos(x)")

	b := chromosomes.NewBuilder()
	b.AddTrait("X")

	optimizef := func(x float64) float64 {
		return ((x * x) + x) * math.Cos(x)
	}

	fitnessf := boundedFitnessFunc1d(min, max, optimizef)

	result := optimize.Optimize(fitnessf, b)

	inputVal := rescale(result.Get("X"), min, max)
	fmt.Println(result, fitnessf(result))

	fmt.Println("Best result: X=", inputVal)
}

func optimize2d() {
	max, min := 2, -2
	fmt.Println("Bounded function optimization in range", min, "to", max)
	fmt.Println("f(x,y) = cos(x)cos(y)e^(-x^2 -y^2) ")

	b := chromosomes.NewBuilder()
	b.AddTrait("X")
	b.AddTrait("Y")

	optimizef := func(x, y float64) float64 {
		return math.Cos(x) * math.Cos(y) * math.Exp(-(x*x)-(y*y))
	}

	fitnessf := boundedFitnessFunc2d(min, max, optimizef)

	result := optimize.Optimize(fitnessf, b)

	inputValX := rescale(result.Get("X"), min, max)
	inputValY := rescale(result.Get("Y"), min, max)
	fmt.Println(result, fitnessf(result))

	fmt.Println("Best result: (X,Y)=", inputValX, inputValY)
}

func main() {
	optimize1d()
	fmt.Println()
	optimize2d()
}
