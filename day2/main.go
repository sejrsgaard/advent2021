package main

import (
	utils "advent/utils"
	_ "embed"
	"fmt"
	"strconv"
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
	fmt.Printf("Part 1 = %d\n", get_position(input))
	fmt.Printf("Part 2 = %d\n", get_position_with_aim(input))

}

func get_position(input string) int {
	horizontal := 0
	depth := 0
	for _, s := range strings.Split(input, "\n") {
		parts := strings.Split(s, " ")
		v := utils.Cast(parts[1])
		if parts[0] == "forward" {
			horizontal = horizontal + v
		}
		if parts[0] == "down" {
			depth = depth + v
		}
		if parts[0] == "up" {
			depth = depth - v
		}
	}
	return horizontal * depth
}

func get_position_with_aim(input string) int {
	horizontal := 0
	depth := 0
	aim := 0
	for _, s := range strings.Split(input, "\n") {
		parts := strings.Split(s, " ")
		v := utils.Cast(parts[1])
		if parts[0] == "forward" {
			horizontal = horizontal + v
			depth = depth + (aim * v)
		}
		if parts[0] == "down" {
			aim = aim + v
		}
		if parts[0] == "up" {
			aim = aim - v
		}
	}
	return horizontal * depth
}

func rate(input string, sw int) int {
	depths := get_depths(input)
	count := 0
	for i := sw; i < len(depths); i++ {
		if depths[i-sw] < depths[i] {
			count++
		}
	}
	return count
}

func get_depths(input string) (depths []int) {
	for _, s := range strings.Split(input, "\n") {
		d, _ := strconv.Atoi(s)
		depths = append(depths, d)
	}
	return depths
}
