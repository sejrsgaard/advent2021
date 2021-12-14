package main

import (
	"advent/utils"
	_ "embed"
	"fmt"
	"regexp"
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

}

func solve_part1(input string) int {
	r := regexp.MustCompile(`(\d+),(\d+) -> (\d+),(\d+)`)
	grid := [][]int{}
	for _, s := range strings.Split(input, "\n") {
		m := r.FindStringSubmatch(s)
		x1 := utils.Cast(m[1])
		y1 := utils.Cast(m[2])
		x2 := utils.Cast(m[3])
		y2 := utils.Cast(m[4])

		max_x := utils.Max(x1, x2)
		max_y := utils.Max(y1, y2)

		for x := len(grid); x < max_x; x++ {
			if x == 0 {
				grid[x] = []int{}
			}
			for y := len(grid[x]); y < max_y; y++ {
				grid[x][y] = 0
			}
		}

		if x1 == x2 {
			if y1 < y2 {
				for y := y1; y < y2; y++ {
					grid[x1][y]++
				}
			} else {
				for y := y2; y < y1; y++ {
					grid[x1][y]++
				}
			}
		}
	}
	fmt.Printf("%#v\n", grid)
	return 0
}
