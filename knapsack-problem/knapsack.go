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

func main() {
	fmt.Println("Knapsack problem")

	b := chromosomes.NewBuilder()
	b.AddTrait("A")

	maxWt := 10
	item0 := knapsackItem{5, 10}
	item1 := knapsackItem{4, 40}
	item2 := knapsackItem{6, 30}
	item3 := knapsackItem{3, 50}

	f := func(c *chromosomes.Chromosome) float64 {
		totalWt := 0
		totalVal := 0
		byte := c.Get("A")

		if hasItemN(byte, 0) {
			totalVal += item0.value
			totalWt += item0.weight
		}
		if hasItemN(byte, 1) {
			totalVal += item1.value
			totalWt += item1.weight
		}
		if hasItemN(byte, 2) {
			totalVal += item2.value
			totalWt += item2.weight
		}
		if hasItemN(byte, 3) {
			totalVal += item3.value
			totalWt += item3.weight
		}

		if totalWt <= maxWt {
			return float64(totalVal)
		} else {
			return float64(-totalWt)
		}
	}

	result := optimize.Optimize(f, b)
	fmt.Println(result, f(result))

	byte := result.Get("A")
	if hasItemN(byte, 0) {
		fmt.Println("contains item0", item0)
	}
	if hasItemN(byte, 1) {
		fmt.Println("contains item1", item1)
	}
	if hasItemN(byte, 2) {
		fmt.Println("contains item2", item2)
	}
	if hasItemN(byte, 3) {
		fmt.Println("contains item3", item3)
	}
}
