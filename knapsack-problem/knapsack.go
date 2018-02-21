package main

import (
	"fmt"
	"github.com/rylans/chromosomes"
	"github.com/rylans/chromosomes/optimize"
)

type knapsackItem struct {
	weight int
	value  int
}

func hasItemN(byte uint8, n int) bool {
	mask := uint8(1 << uint(n))
	return (byte & mask) > 0
}

func itemFitnessFn(maxWt int, items []knapsackItem) func(c *chromosomes.Chromosome) float64 {
	fitnessFn := func(c *chromosomes.Chromosome) float64 {
		totalWt, totalVal := 0, 0
		byte := c.Get("A")
		byte2 := c.Get("B")

		for i := 0; i < len(items); i++ {
			if i > 8 {
				byte = byte2
			}
			if hasItemN(byte, i) {
				totalVal += items[i%8].value
				totalWt += items[i%8].weight
			}
		}

		if totalWt <= maxWt {
			return float64(totalVal)
		} else {
			return float64(-totalWt)
		}
	}

	return fitnessFn
}

func knapsack4() {
	fmt.Println("Knapsack problem with four items. Max weight: 10")

	b := chromosomes.NewBuilder()
	b.AddTrait("A")
	b.AddTrait("B")

	maxWt := 10
	items := make([]knapsackItem, 0)

	items = append(items, knapsackItem{5, 10})
	items = append(items, knapsackItem{4, 40})
	items = append(items, knapsackItem{6, 30})
	items = append(items, knapsackItem{3, 50})

	f := itemFitnessFn(maxWt, items)
	result := optimize.Optimize(f, b)
	fmt.Println("Result:", result, f(result)) // Best value: 90

	byte := result.Get("A")
	for i := 0; i < len(items); i++ {
		if hasItemN(byte, i) {
			fmt.Println("Contains item #", i, items[i])
		}
	}
}

// Example comes from
// G. Zäpfel et al. Metaheuristic Search Concepts: A Tutorial with Applications to Production and Logistics
func knapsack10() {
	fmt.Println("Knapsack problem with ten items. Max weight: 113")

	b := chromosomes.NewBuilder()
	b.AddTrait("A")
	b.AddTrait("B")

	maxWt := 113
	items := make([]knapsackItem, 0)

	items = append(items, knapsackItem{32, 727})
	items = append(items, knapsackItem{40, 763})
	items = append(items, knapsackItem{44, 60})
	items = append(items, knapsackItem{20, 606})
	items = append(items, knapsackItem{1, 45})
	items = append(items, knapsackItem{29, 370})
	items = append(items, knapsackItem{3, 414})
	items = append(items, knapsackItem{13, 880})
	items = append(items, knapsackItem{6, 133})
	items = append(items, knapsackItem{39, 820})

	f := itemFitnessFn(maxWt, items)
	result := optimize.Optimize(f, b)
	fmt.Println("Result:", result, f(result)) // Best value: 3580

	byte := result.Get("A")
	byte2 := result.Get("B")
	for i := 0; i < len(items); i++ {
		if i > 8 {
			byte = byte2
		}
		if hasItemN(byte, i) {
			fmt.Println("Contains item #", i, items[i%8])
		}
	}
}

func main() {
	knapsack4()
	fmt.Println()
	knapsack10()
}
