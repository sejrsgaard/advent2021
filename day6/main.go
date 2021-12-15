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
	result := solve(input, 80)
	fmt.Printf("Part 1 = %d\n", result)
	result = solve(input, 256)
	fmt.Printf("Part 2 = %d\n", result)
}

func solve(input string, days int) int64 {
	fish := make(map[int]int64)
	for _, s := range strings.Split(input, ",") {
		fish[utils.Cast(s)]++
	}

	age := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	for day := 0; day < days; day++ {
		new_fish := make(map[int]int64)
		for k := range age {
			if k == 0 {
				new_fish[8] = fish[k]
				new_fish[6] = fish[k]
			} else if k == 7 {
				new_fish[6] += fish[k]
			} else {
				new_fish[k-1] = fish[k]
			}
		}
		fish = new_fish
	}

	sum := int64(0)
	for _, v := range fish {
		sum += v
	}

	return sum
}
