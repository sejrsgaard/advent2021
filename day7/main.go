package main

import (
	"advent/utils"
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func init() {
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
}

func main() {
	result := solve(input)
	fmt.Printf("Part 1 = %d\n", result)
	result = solve2(input)
	fmt.Printf("Part 2 = %d\n", result)
}

func solve(input string) int {
	positions := []int{}
	for _, s := range strings.Split(input, ",") {
		positions = append(positions, utils.Cast(s))
	}

	// Find max
	max := 0
	for _, v := range positions {
		if v > max {
			max = v
		}
	}

	fuel := 2147483647
	for i := 0; i <= max; i++ {
		f := 0
		for _, v := range positions {
			f += utils.Abs(i - v)
		}
		if f < fuel {
			fuel = f
		}
	}

	return fuel
}

func solve2(input string) int {
	positions := []int{}
	for _, s := range strings.Split(input, ",") {
		positions = append(positions, utils.Cast(s))
	}

	// Find max
	max := 0
	for _, v := range positions {
		if v > max {
			max = v
		}
	}

	fuel := 2147483647
	for i := 0; i <= max; i++ {
		f := 0
		for _, v := range positions {
			n := utils.Abs(i - v)
			f += (n * (n + 1)) / 2
		}
		if f < fuel {
			fuel = f
		}
	}

	return fuel
}
