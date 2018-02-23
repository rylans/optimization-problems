// One-dimensional function optimization
package main

import (
	"math"
	"fmt"
	"github.com/rylans/chromosomes"
	"github.com/rylans/chromosomes/optimize"
)

func rescale(val uint8, min int, max int) float64 {
  newRange := float64(max) - float64(min)
  return (newRange/255.0)*(float64(val) - 255.0) + float64(max)
}

type optimizeFn func(x float64) float64

func boundedFitnessFunc(min int, max int, optimizef optimizeFn) func(c *chromosomes.Chromosome) float64{ 
  f := func(c *chromosomes.Chromosome) float64 {
    realVal := rescale(c.Get("X"), min, max)
    return optimizef(realVal)
  }
  return f
}

func main (){
  max, min := 10, -10
  fmt.Println("Bounded function optimization in range", min, "to", max)
  fmt.Println("f(x) = (x^2 + x) * cos(x)")

  b := chromosomes.NewBuilder()
  b.AddTrait("X")

  optimizef := func(x float64) float64{
    return ((x*x) + x) * math.Cos(x)
  }

  fitnessf := boundedFitnessFunc(min, max, optimizef)

  result := optimize.Optimize(fitnessf, b)

  inputVal := rescale(result.Get("X"), min, max)
  fmt.Println(result, fitnessf(result))

  fmt.Println("Best result: X=",inputVal)
}
