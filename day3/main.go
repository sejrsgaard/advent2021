package main

import (
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
	result := solve_part1(input)
	fmt.Printf("%d\n", result)
	result = solve_part2(input)
	fmt.Printf("%d\n", result)
}

func solve_part1(input string) int {
	bits := []int{}
	for n, s := range strings.Split(input, "\n") {
		if n == 0 {
			bits = make([]int, len(s))
		}
		for i, c := range s {
			if c == '0' {
				bits[i] = bits[i] - 1
			} else {

				bits[i] = bits[i] + 1
			}
		}

	}

	gamma := ""
	epsilon := ""
	for _, v := range bits {
		if v >= 0 {
			gamma += "1"
			epsilon += "0"
		} else {
			gamma += "0"
			epsilon += "1"
		}
	}

	g, _ := strconv.ParseInt(gamma, 2, 64)
	e, _ := strconv.ParseInt(epsilon, 2, 64)
	return int(g * e)
}

func solve_part2(input string) int {
	entries := strings.Split(input, "\n")

	i := 0
	for {
		keep := []string{}
		v := 0
		for _, s := range entries {
			c := s[i]
			if c == '0' {
				v--
			} else {
				v++
			}
		}

		for _, s := range entries {
			if v >= 0 {
				if s[i] == '1' {
					keep = append(keep, s)
				}
			} else {
				if s[i] == '0' {
					keep = append(keep, s)
				}
			}
		}

		entries = keep
		i++
		if len(entries) <= 1 {
			break
		}
	}
	oxygen := entries[0]

	entries = strings.Split(input, "\n")
	i = 0
	for {
		keep := []string{}
		v := 0
		for _, s := range entries {
			c := s[i]
			if c == '0' {
				v--
			} else {
				v++
			}
		}

		for _, s := range entries {
			if v >= 0 {
				if s[i] == '0' {
					keep = append(keep, s)
				}
			} else {
				if s[i] == '1' {
					keep = append(keep, s)
				}
			}
		}

		entries = keep
		i++
		if len(entries) <= 1 {
			break
		}
	}
	co2 := entries[0]

	o, _ := strconv.ParseInt(oxygen, 2, 64)
	c, _ := strconv.ParseInt(co2, 2, 64)
	return int(o * c)
}
