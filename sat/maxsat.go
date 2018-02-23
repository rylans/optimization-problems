package main

import (
	"fmt"
	"github.com/rylans/chromosomes"
	"github.com/rylans/chromosomes/optimize"
)

type predicateFn func(vars []bool) bool

func bitstringToBoolSlice(bits uint8) []bool {
	bools := make([]bool, 0)
	for i := uint8(0); i < 8; i++ {
		if (bits & (1 << i)) > 0 {
			bools = append(bools, true)
		} else {
			bools = append(bools, false)
		}
	}
	return bools
}

// returns a function that returns the number of predicates that c satisfies
func fitnessProvider(predicates []predicateFn) func(c *chromosomes.Chromosome) float64 {
	f := func(c *chromosomes.Chromosome) float64 {
		satisfied := float64(0)

		bits := c.Get("A")
		for _, p := range predicates {
			if p(bitstringToBoolSlice(bits)) {
				satisfied += 1
			}
		}

		return satisfied
	}
	return f
}

func main() {
	fmt.Println("Maximum satisfiability problem (MAX-SAT)")

	b := chromosomes.NewBuilder()
	b.AddTrait("A")

	predicates := make([]predicateFn, 0)
	p1 := func(vars []bool) bool {
		return vars[0] && vars[1] && !vars[2]
	}
	p2 := func(vars []bool) bool {
		return !vars[0] || vars[3]
	}
	p3 := func(vars []bool) bool {
		return !vars[3] && vars[4]
	}
	p4 := func(vars []bool) bool {
		return vars[5] && !vars[6] && vars[7]
	}
	predicates = append(predicates, p1)
	predicates = append(predicates, p2)
	predicates = append(predicates, p3)
	predicates = append(predicates, p4)

	result := optimize.Optimize(fitnessProvider(predicates), b)

	satisfied := fitnessProvider(predicates)(result)
	fmt.Println(result, satisfied)
	fmt.Println(bitstringToBoolSlice(result.Get("A")))
	fmt.Println("Out of", len(predicates), "predicates,", satisfied, "are satisfied.")
}
