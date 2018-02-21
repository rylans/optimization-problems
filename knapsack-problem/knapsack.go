package main

import (
  "fmt"
  "github.com/rylans/chromosomes"
  "github.com/rylans/chromosomes/optimize"
)

type knapsackItem struct {
  weight int
  value int
}

func main(){
  fmt.Println("Knapsack problem")

  b := chromosomes.NewBuilder()
  b.AddTrait("A")

  f := func(c *chromosomes.Chromosome) float64 {
    return 0.0
  }

  result := optimize.Optimize(f, b)
  fmt.Println(result)

}
