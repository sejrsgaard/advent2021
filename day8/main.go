package main

import (
	"advent/utils"
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
	result := solve(input)
	fmt.Printf("Part 1 = %d\n", result)
	result = solve2(input)
	fmt.Printf("Part 2 = %d\n", result)
}

func solve(input string) int {
	io := [][2][]string{}
	for _, s := range strings.Split(input, "\n") {
		parts := strings.Split(s, " | ")
		input := strings.Fields(parts[0])
		output := strings.Fields(parts[1])
		io = append(io, [2][]string{input, output})
	}

	result := 0
	for _, v := range io {
		for _, s := range v[1] {
			l := len(s)
			if l == 2 || l == 3 || l == 4 || l == 7 {
				result++
			}
		}
	}
	return result
}

/*
 0 0 0
5      1
5	   1
5	   1
 6 6 6
4	   2
4	   2
4	   2
 3 3 3
*/
func solve2(input string) (result int) {
	io := [][2][]string{}
	for _, s := range strings.Split(input, "\n") {
		parts := strings.Split(s, " | ")
		input := strings.Split(parts[0], " ")
		output := strings.Split(parts[1], " ")
		io = append(io, [2][]string{input, output})
	}

	for _, v := range io {
		translation := make(map[string]string)
		numbers := make(map[int]string, 10)

		digits := []string{}
		for _, s := range v[0] {
			digits = append(digits, utils.Sort(s))
		}

		for _, s := range digits {
			if len(s) == 2 {
				numbers[1] = s
				translation[s] = "1"
			} else if len(s) == 3 {
				numbers[7] = s
				translation[s] = "7"
			} else if len(s) == 4 {
				numbers[4] = s
				translation[s] = "4"
			} else if len(s) == 7 {
				numbers[8] = s
				translation[s] = "8"
			}
		}

		segments := [7]rune{}

		// Find segment 0
		for _, r := range numbers[7] {
			if !strings.ContainsRune(numbers[1], r) {
				segments[0] = r
			}
		}

		// Find segment 3
		for _, s := range digits {
			if len(s) == 5 && utils.ContainsRunes(s, []rune(numbers[1])) {
				translation[s] = "3"
				numbers[3] = s
				segments[3] = except(s, numbers[4], numbers[7])
				break
			}
		}

		// Find segment 6
		segments[6] = except(numbers[3], numbers[1], string([]rune{segments[0], segments[3]}))

		// Find segment 5
		for _, s := range digits {
			if len(s) == 6 && utils.ContainsRunes(s, utils.Runes(numbers[4])) {
				translation[s] = "9"
				numbers[9] = s
				segments[5] = except(s, numbers[1], string([]rune{segments[0], segments[3], segments[6]}))
				break
			}
		}

		// Find segment 1, 2
		for _, s := range digits {
			if len(s) == 5 && strings.ContainsRune(s, segments[5]) {
				translation[s] = "5"
				numbers[5] = s
				segments[1] = except(numbers[1], numbers[5])
				segments[2] = except(numbers[1], string([]rune{segments[1]}))
			}
		}

		// Find segment 4
		segments[4] = except(numbers[8], numbers[9])

		for _, s := range digits {
			if len(s) == 5 && s != numbers[3] && s != numbers[5] {
				translation[s] = "2"
			} else if len(s) == 6 && s != numbers[9] {
				if strings.ContainsRune(s, segments[6]) {
					translation[s] = "6"
				} else {
					translation[s] = "0"
				}
			}
		}

		out := ""
		for _, s := range v[1] {
			s = utils.Sort(s)
			out += translation[s]
		}
		i, _ := strconv.Atoi(out)
		result += i
	}

	return result
}

func except(a string, b ...string) rune {
	var out string
	for _, r := range a {
		found := false
		for _, s := range b {
			for _, rr := range s {
				if r == rr {
					found = true
					break
				}
			}
			if found {
				break
			}
		}
		if !found {
			out += string(r)
		}
	}
	if len(out) != 1 {
		panic("expected 1 rune")
	}

	return rune(out[0])
}
