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
	result := solve_part1(input)
	fmt.Printf("Part 1 = %d\n", result)
	result = solve_part2(input)
	fmt.Printf("Part 2 = %d\n", result)
}

func solve_part1(input string) int {
	r := regexp.MustCompile(`(\d+),(\d+) -> (\d+),(\d+)`)
	lines := [][]int{}
	for _, s := range strings.Split(input, "\n") {
		m := r.FindStringSubmatch(s)
		x1 := utils.Cast(m[1])
		y1 := utils.Cast(m[2])
		x2 := utils.Cast(m[3])
		y2 := utils.Cast(m[4])

		lines = append(lines, []int{x1, y1, x2, y2})
	}

	max_x := 0
	max_y := 0
	for _, l := range lines {
		max_x = utils.Max(max_x, utils.Max(l[0], l[2]))
		max_y = utils.Max(max_y, utils.Max(l[1], l[3]))
	}

	grid := make([][]int, max_x+1)
	for i := range grid {
		grid[i] = make([]int, max_y+1)
	}

	for _, l := range lines {

		if l[0] == l[2] {
			if l[1] < l[3] {
				for y := l[1]; y <= l[3]; y++ {
					grid[l[0]][y]++
				}
			} else {
				for y := l[3]; y <= l[1]; y++ {
					grid[l[0]][y]++
				}
			}
		}
		if l[1] == l[3] {
			if l[0] < l[2] {
				for x := l[0]; x <= l[2]; x++ {
					grid[x][l[1]]++
				}
			} else {
				for x := l[2]; x <= l[0]; x++ {
					grid[x][l[1]]++
				}
			}
		}
	}

	count := 0
	for _, row := range grid {
		for _, n := range row {
			if n >= 2 {
				count++
			}
		}
	}

	return count
}

func solve_part2(input string) int {
	r := regexp.MustCompile(`(\d+),(\d+) -> (\d+),(\d+)`)
	lines := [][]int{}
	for _, s := range strings.Split(input, "\n") {
		m := r.FindStringSubmatch(s)
		x1 := utils.Cast(m[1])
		y1 := utils.Cast(m[2])
		x2 := utils.Cast(m[3])
		y2 := utils.Cast(m[4])

		lines = append(lines, []int{x1, y1, x2, y2})
	}

	max_x := 0
	max_y := 0
	for _, l := range lines {
		max_x = utils.Max(max_x, utils.Max(l[0], l[2]))
		max_y = utils.Max(max_y, utils.Max(l[1], l[3]))
	}

	grid := make([][]int, max_x+1)
	for i := range grid {
		grid[i] = make([]int, max_y+1)
	}

	for _, l := range lines {
		x1, x2 := l[0], l[2]
		y1, y2 := l[1], l[3]

		if x1 == x2 {
			n, m := from_to(y1, y2)
			for i := n; i <= m; i++ {
				grid[x1][i]++
			}
		} else if y1 == y2 {
			n, m := from_to(x1, x2)
			for i := n; i <= m; i++ {
				grid[i][y1]++
			}
		} else {
			it := utils.Abs(x2 - x1)
			for i := 0; i <= it; i++ {
				if x2 > x1 {
					if y2 > y1 {
						grid[x1+i][y1+i]++
					} else {
						grid[x1+i][y1-i]++
					}
				} else {
					if y2 > y1 {
						grid[x1-i][y1+i]++
					} else {
						grid[x1-i][y1-i]++
					}
				}
			}
		}
	}

	count := 0
	for _, row := range grid {
		for _, n := range row {
			if n >= 2 {
				count++
			}
		}
	}

	return count
}

func from_to(x1 int, x2 int) (int, int) {
	if x2 > x1 {
		return x1, x2
	}
	return x2, x1
}
