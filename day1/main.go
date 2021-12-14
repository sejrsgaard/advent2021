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
	sw1 := rate(input, 1)
	fmt.Printf("%d\n", sw1)
	sw3 := rate(input, 3)
	fmt.Printf("%d\n", sw3)
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
